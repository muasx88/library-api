package book

import "time"

type BookEntity struct {
	Id              int       `db:"id"`
	Title           string    `db:"title"`
	Author          string    `db:"author"`
	PublicationYear int       `db:"publication_year"`
	Genre           string    `db:"genre"`
	ISBN            string    `db:"isbn"`
	Stock           int       `db:"stock"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

func NewFromAddBookRequestPayload(req AddBookRequestPayload) BookEntity {
	return BookEntity{
		Title:           req.Title,
		Author:          req.Author,
		PublicationYear: req.PublicationYear,
		Genre:           req.Genre,
		ISBN:            req.ISBN,
		Stock:           req.Stock,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (book BookEntity) IsExists() bool {

	return book.Id != 0
}
