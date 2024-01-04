package book

import "time"

type BookResponse struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationYear int       `json:"publication_year"`
	Genre           string    `json:"genre"`
	ISBN            string    `json:"isbn"`
	Stock           int       `json:"stock"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func NewBookResponse(book BookEntity) (res BookResponse) {
	return BookResponse{
		Id:              book.Id,
		Title:           book.Title,
		Author:          book.Author,
		PublicationYear: book.PublicationYear,
		Genre:           book.Genre,
		ISBN:            book.ISBN,
		Stock:           book.Stock,
		CreatedAt:       book.CreatedAt,
		UpdatedAt:       book.UpdatedAt,
	}
}

func NewBooksResponse(books []BookEntity) (res []BookResponse) {
	for _, book := range books {
		newBook := NewBookResponse(book)
		res = append(res, newBook)
	}

	return
}
