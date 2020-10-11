package book

//Book A single book
type Book struct {
	Title          string  `json:"title"`
	ID             int     `json:"book_id"`
	ImageURL       string  `json:"image_url"`
	MiniImageURL   string  `json:"smaill_image_url"`
	Ratings1       int     `json:"ratings_1"`
	Ratings2       int     `json:"ratings_2"`
	Ratings3       int     `json:"ratings_3"`
	Ratings4       int     `json:"ratings_4"`
	Ratings5       int     `json:"ratings_5"`
	AverageRating  float32 `json:"average_rating"`
	RatingCount    int     `json:"rating_count"`
	Authors        string  `json:"authors"`
	AvailableCount int     `json:"books_count"`
}

//Books a slice of Book
type Books []Book
