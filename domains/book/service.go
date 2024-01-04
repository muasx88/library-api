package book

import (
	"context"

	"github.com/muasx88/library-api/internal/response"
)

type IService interface {
	GetAll(context.Context) ([]BookEntity, error)
	DetailById(context.Context, int) (BookEntity, error)
	DetailByIsbn(context.Context, string) (BookEntity, error)
	Store(context.Context, AddBookRequestPayload) error
}

type service struct {
	repo IRepository
}

func newService(repo IRepository) IService {
	return &service{
		repo: repo,
	}
}

func (s service) GetAll(ctx context.Context) ([]BookEntity, error) {
	return s.repo.GetAll(ctx)
}

func (s service) DetailById(ctx context.Context, id int) (BookEntity, error) {
	return s.repo.DetailById(ctx, id)
}

func (s service) DetailByIsbn(ctx context.Context, isbn string) (BookEntity, error) {
	return s.repo.DetailByIsbn(ctx, isbn)
}

func (s service) Store(ctx context.Context, req AddBookRequestPayload) (err error) {
	entity := NewFromAddBookRequestPayload(req)

	model, err := s.repo.DetailByIsbn(ctx, entity.ISBN)
	if err != nil {
		if err != response.ErrBookNotFound {
			return
		}
	}

	if model.IsExists() {
		return response.ErrBookAlreadyExists
	}

	return s.repo.Store(ctx, entity)
}
