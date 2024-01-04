package auth

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/muasx88/library-api/internal/response"
)

type IRepository interface {
	GetAuthByEmail(ctx context.Context, email string) (entity UserEntity, err error)
	CreateAuth(ctx context.Context, entity UserEntity) (err error)
}

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) IRepository {
	return &repository{
		db: db,
	}
}

func (r repository) CreateAuth(ctx context.Context, entity UserEntity) (err error) {
	query := `
		INSERT INTO users (email, password, role, name) 
		VALUES (:email, :password, :role, :name)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error prepare statement db. %s", err.Error())
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, entity)

	return
}

// GetAuthByEmail implements Repository.
func (r repository) GetAuthByEmail(ctx context.Context, email string) (entity UserEntity, err error) {
	query := `
		SELECT id, email, password, role, name, created_at, updated_at
		FROM users
		WHERE email = ?
	`

	err = r.db.GetContext(ctx, &entity, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrUserNotFound
			return
		}
		return
	}

	return
}
