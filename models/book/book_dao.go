package book

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/pt-abhishek/bookstore-api/databases/elasticsearch"
	"github.com/pt-abhishek/bookstore-api/databases/mysql"
	"github.com/pt-abhishek/bookstore-api/utils/errors"
	"github.com/tidwall/gjson"
)

//BookDAO is the dao
var (
	BookDAO                BooksDAOInterface = &booksDAO{}
	queryFindAllPagination                   = "SELECT book_id, books_count, authors, title, image_url, small_image_url, average_rating, ratings_count, ratings_1, ratings_2, ratings_3, ratings_4, ratings_5 FROM books LIMIT ?,?"
	booksESIndex                             = "books_sql"
)

//BooksDAOInterface is the books dao
type BooksDAOInterface interface {
	SearchByName(name string) (Books, *errors.RestErr)
	GetAll(page int64, pageSize int64) (Books, *errors.RestErr)
}

type booksDAO struct{}

func (dao *booksDAO) SearchByName(name string) (Books, *errors.RestErr) {
	client := elasticsearch.ElasticClient.GetClient()
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": name,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(booksESIndex),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	defer res.Body.Close()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	buf.ReadFrom(res.Body)
	books := gjson.GetManyBytes(buf.Bytes(), "hits.hits.#._source")
	results := make([]Book, 0)
	for _, book := range books[0].Array() {
		bookJSON := book.Raw
		var b Book
		if err := json.Unmarshal([]byte(bookJSON), &b); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, b)
	}
	return results, nil
}

func (dao *booksDAO) GetAll(page int64, pageSize int64) (Books, *errors.RestErr) {
	client := mysql.SQLClient.GetClient()
	stmt, err := client.Prepare(queryFindAllPagination)
	if err != nil {
		log.Fatal(err.Error())
		return nil, errors.NewInternalServerError("Error creating new query")
	}
	defer stmt.Close()
	rows, err := stmt.Query((page-1)*pageSize, pageSize)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	results := make([]Book, 0)
	for rows.Next() {
		var book Book
		if err = rows.Scan(
			&book.ID,
			&book.AvailableCount,
			&book.Authors,
			&book.Title,
			&book.ImageURL,
			&book.MiniImageURL,
			&book.AverageRating,
			&book.RatingCount,
			&book.Ratings1,
			&book.Ratings2,
			&book.Ratings3,
			&book.Ratings4,
			&book.Ratings5); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, book)
	}
	return results, nil
}
