package transaction

import "github.com/muasx88/library-api/internal/response"

type BookEntity struct {
	Id              int    `db:"id"`
	Title           string `db:"title"`
	Author          string `db:"author"`
	PublicationYear int    `db:"publication_year"`
	Genre           string `db:"genre"`
	ISBN            string `db:"isbn"`
	Stock           int    `db:"stock"`
}

func (b *BookEntity) IsExists() bool {
	return b.Id != 0
}

func (b *BookEntity) DecrementStockBook() (err error) {
	if b.Stock == 0 {
		return response.ErrBookStock
	}

	b.Stock--
	return
}

func (b *BookEntity) IncrementStockBook() {
	b.Stock++
}
