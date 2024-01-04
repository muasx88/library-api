package transaction

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/muasx88/library-api/internal/response"
)

type IRepository interface {
	repositoryDb
	repositoryTransaction
	repositoryBook
}

type repositoryDb interface {
	Begin(context.Context) (*sqlx.Tx, error)
	Rollback(context.Context, *sqlx.Tx) error
	Commit(context.Context, *sqlx.Tx) error
}

type repositoryTransaction interface {
	StoreTransaction(context.Context, *sqlx.Tx, TransactionEntity) error
	UpdateTransaction(context.Context, *sqlx.Tx, TransactionEntity) error
	DetailTransaction(context.Context, int) (TransactionEntity, error)
	UserHistoryTransaction(context.Context, int) ([]TransactionEntity, error)
	UserLatestTransaction(context.Context, int) (TransactionEntity, error)
}

type repositoryBook interface {
	DetailBookByIsbn(context.Context, string) (BookEntity, error)
	UpdateBook(context.Context, *sqlx.Tx, BookEntity) error
}

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) IRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Begin(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = r.db.BeginTxx(ctx, &sql.TxOptions{})
	return
}

func (r *repository) Commit(ctx context.Context, tx *sqlx.Tx) (err error) {
	err = tx.Commit()
	return
}

func (r *repository) Rollback(ctx context.Context, tx *sqlx.Tx) (err error) {
	err = tx.Rollback()
	return
}

func (r *repository) UserHistoryTransaction(ctx context.Context, userId int) (res []TransactionEntity, err error) {
	query := `
		SELECT t.id, t.user_id, t.book_id, b.title as book_title, b.isbn as book_isbn,
		t.duration, t.start_date, t.end_date, t.return_date, t.status, t.created_at, t.updated_at
		FROM transactions t
		LEFT JOIN books b on t.book_id = b.id
		WHERE t.user_id = ?
		ORDER BY t.created_at DESC
	`

	err = r.db.SelectContext(ctx, &res, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrTransactionNotFound
		}

		return
	}

	return
}

func (r *repository) UserLatestTransaction(ctx context.Context, userId int) (entity TransactionEntity, err error) {
	query := `
		SELECT t.id, t.user_id, t.book_id, b.title as book_title, b.isbn as book_isbn,
		t.duration, t.start_date, t.end_date, t.return_date, t.status, t.created_at, t.updated_at
		FROM transactions t
		LEFT JOIN books b on t.book_id = b.id
		WHERE t.user_id = ?
		ORDER BY t.created_at DESC
	`

	err = r.db.GetContext(ctx, &entity, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrTransactionNotFound
			return
		}
		return
	}

	return
}

// DetailTransaction implements IRepositori.
func (r *repository) DetailTransaction(ctx context.Context, id int) (entity TransactionEntity, err error) {
	query := `
		SELECT t.id, t.user_id, u.name as user_name, t.book_id, b.title as book_title, b.isbn as book_isbn,
		t.duration, t.start_date, t.end_date, t.return_date, t.status, t.created_at, t.updated_at
		FROM transactions t
		LEF JOIN users u on t.user_id = u.id
		LEF JOIN books b on t.book_id = b.id
		WHERE t.id = ?
	`

	err = r.db.GetContext(ctx, &entity, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrTransactionNotFound
			return
		}
		return
	}

	return
}

// StoreTransaction implements IRepositori.
func (r *repository) StoreTransaction(ctx context.Context, tx *sqlx.Tx, entity TransactionEntity) (err error) {
	query := `
		INSERT into transactions (user_id, book_id, duration, start_date, end_date)
		VALUES (:user_id, :book_id, :duration, :start_date, :end_date)
	`

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		err = fmt.Errorf("error prepare statement db. %s", err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, entity)
	if err != nil {
		err = fmt.Errorf("error save transaction into db. %s", err.Error())
		return
	}

	// lastId, err := result.LastInsertId()
	// if err != nil {
	// 	err = fmt.Errorf("error get last inserted id transaction %s", err.Error())
	// 	return
	// }

	// return r.Detail(ctx, int(lastId))

	return
}

// StoreTransaction implements IRepositori.
func (r *repository) UpdateTransaction(ctx context.Context, tx *sqlx.Tx, entity TransactionEntity) (err error) {
	query := `
		UPDATE transactions SET return_date = :return_date, is_late = :is_late
		WHERE id = :id
	`

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		err = fmt.Errorf("error prepare statement db. %s", err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, entity)
	if err != nil {
		err = fmt.Errorf("error update transaction into db. %s", err.Error())
		return
	}

	return
}

func (r *repository) DetailBookByIsbn(ctx context.Context, isbn string) (entity BookEntity, err error) {
	query := `
		SELECT id, title, author, publication_year, genre, isbn, stock
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
func (r *repository) UpdateBook(ctx context.Context, tx *sqlx.Tx, entity BookEntity) (err error) {
	query := `
		UPDATE books SET stock = :stock
		WHERE id = :id
	`

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error prepare statement db. %s", err.Error())
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, entity)
	if err != nil {
		return
	}

	return
}
