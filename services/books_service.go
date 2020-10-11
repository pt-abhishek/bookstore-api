package services

import (
	"github.com/pt-abhishek/bookstore-api/models/book"
	"github.com/pt-abhishek/bookstore-api/utils/errors"
)

type bookService struct{}

//BookServiceInterface is instantiated
type BookServiceInterface interface {
	GetBySearchText(searchText string) (book.Books, *errors.RestErr)
	GetAllWithPagination(page int64, pageSize int64) (book.Books, *errors.RestErr)
}

//BookService An instance of the BooksService
var BookService BookServiceInterface = &bookService{}

func (s *bookService) GetBySearchText(searchText string) (book.Books, *errors.RestErr) {
	books, err := book.BookDAO.SearchByName(searchText)
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, errors.NewResourceNotFoundError("No books found")
	}
	return books, nil
}

func (s *bookService) GetAllWithPagination(page int64, pageSize int64) (book.Books, *errors.RestErr) {
	books, err := book.BookDAO.GetAll(page, pageSize)
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, errors.NewResourceNotFoundError("No books found")
	}
	return books, nil
}
