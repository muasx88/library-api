package transaction

import (
	"context"
	"encoding/json"
	"log"

	"github.com/muasx88/library-api/internal/response"
)

type IService interface {
	Borrow(context.Context, BorrowRequest) error
	Return(context.Context, ReturnRequest) error
	UserHistoryTransaction(context.Context, int) ([]TransactionEntity, error)
}

type service struct {
	repo IRepository
}

func newService(repo IRepository) IService {
	return &service{
		repo: repo,
	}
}

func (s *service) UserHistoryTransaction(ctx context.Context, userId int) ([]TransactionEntity, error) {
	return s.repo.UserHistoryTransaction(ctx, userId)
}

// Borrow implements IService.
func (s *service) Borrow(ctx context.Context, req BorrowRequest) (err error) {
	var availBorrow bool
	availBorrow = true

	book, err := s.repo.DetailBookByIsbn(ctx, req.BookIsbn)
	if err != nil {
		return
	}

	latestTrx, err := s.repo.UserLatestTransaction(ctx, req.UserId)
	if err != nil {
		if err != response.ErrTransactionNotFound {
			return err
		}
	}

	// check if latest transaction still borrowing
	if latestTrx.Id != 0 && latestTrx.Status != TransactionStatus_Done {
		availBorrow = false
	}

	if !availBorrow {
		return response.ErrForbiddenTransaction
	}

	txEntity := NewBorrowRequest(req, book)

	tx, err := s.repo.Begin(ctx)
	if err != nil {
		return
	}

	defer s.repo.Rollback(ctx, tx)

	if err = s.repo.StoreTransaction(ctx, tx, txEntity); err != nil {
		return
	}

	// decrease book stock
	if err = book.DecrementStockBook(); err != nil {
		return err
	}

	// update current stock
	if err = s.repo.UpdateBook(ctx, tx, book); err != nil {
		return
	}

	if err = s.repo.Commit(ctx, tx); err != nil {
		return
	}

	return
}

// Return implements IService.
func (s *service) Return(ctx context.Context, req ReturnRequest) (err error) {

	// check user hostory transaction
	var avaibleToReturn bool
	avaibleToReturn = true

	latestTrx, err := s.repo.UserLatestTransaction(ctx, req.UserId)

	jsonTrx, _ := json.Marshal(latestTrx)
	log.Println(string(jsonTrx))

	if err != nil {
		if err != response.ErrTransactionNotFound {
			avaibleToReturn = false
		} else {
			return
		}
	}

	if latestTrx.Id != 0 && latestTrx.Status != TransactionStatus_Borrowing {
		avaibleToReturn = false
	}

	if !avaibleToReturn {
		err = response.ErrForbiddenReturnTransaction
		return
	}

	// check if the returning book is the same as the latest borrowing book
	if latestTrx.BookIsbn != req.BookIsbn {
		return response.ErrWrongReturnTransaction
	}

	log.Println(latestTrx.BookIsbn)
	log.Println(req.BookIsbn)

	tx, err := s.repo.Begin(ctx)
	if err != nil {
		return
	}

	defer s.repo.Rollback(ctx, tx)

	txEntity := NewReturnRequest(req, latestTrx)
	txEntity.SetIsLate() // check if late for return

	if err = s.repo.UpdateTransaction(ctx, tx, txEntity); err != nil {
		return
	}

	book, err := s.repo.DetailBookByIsbn(ctx, latestTrx.BookIsbn)
	if err != nil {
		return err
	}

	book.IncrementStockBook() // increase / increment stock after returning the book

	// update current stock
	if err = s.repo.UpdateBook(ctx, tx, book); err != nil {
		return
	}

	if err = s.repo.Commit(ctx, tx); err != nil {
		return
	}

	return
}
