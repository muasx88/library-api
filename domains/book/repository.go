package book

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/muasx88/library-api/internal/response"
)

type IRepository interface {
	GetAll(context.Context) ([]BookEntity, error)
	Store(context.Context, BookEntity) error
	DetailById(context.Context, int) (BookEntity, error)
	DetailByIsbn(context.Context, string) (BookEntity, error)
}

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) IRepository {
	return &repository{
		db: db,
	}
}

func (r repository) GetAll(ctx context.Context) (res []BookEntity, err error) {

	query := `
		SELECT id, title, author, publication_year, genre, isbn, stock, created_at, updated_at
		FROM books
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	var book BookEntity
	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.PublicationYear, &book.Genre, &book.ISBN, &book.Stock, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scan db %s", err.Error())
		}

		res = append(res, book)
	}

	log.Println("Success get data")

	return
}

func (r repository) Store(ctx context.Context, entity BookEntity) (err error) {
	query := `
		INSERT INTO books (title, author, publication_year, genre, isbn, stock) 
		VALUES (:title, :author, :publication_year, :genre, :isbn, :stock)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error prepare statement db. %s", err.Error())
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, entity)

	return
}

func (r repository) DetailById(ctx context.Context, id int) (entity BookEntity, err error) {
	query := `
		SELECT id, title, author, publication_year, genre, isbn, stock, created_at, updated_at
		FROM books
		WHERE id = ?
	`

	err = r.db.GetContext(ctx, &entity, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrBookNotFound
			return
		}
		return
	}

	return

}
func (r repository) DetailByIsbn(ctx context.Context, isbn string) (entity BookEntity, err error) {
	query := `
		SELECT id, title, author, publication_year, genre, isbn, stock, created_at, updated_at
		FROM books
		WHERE isbn = ?
	`

	err = r.db.GetContext(ctx, &entity, query, isbn)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrBookNotFound
			return
		}
		return
	}

	return
}
