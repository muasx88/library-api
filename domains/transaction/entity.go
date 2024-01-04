package transaction

import (
	"time"
)

type TransactionStatus string

const (
	TransactionStatus_Borrowing TransactionStatus = "BORROWING"
	TransactionStatus_Done      TransactionStatus = "DONE"
)

type TransactionEntity struct {
	Id         int               `db:"id"`
	UserId     int               `db:"user_id"`
	BookId     int               `db:"book_id"`
	BookTitle  string            `db:"book_title"`
	BookIsbn   string            `db:"book_isbn"`
	Duration   int               `db:"duration"`
	StartDate  time.Time         `db:"start_date"`
	EndDate    time.Time         `db:"end_date"`
	ReturnDate *time.Time        `db:"return_date"`
	Status     TransactionStatus `db:"status"`
	IsLate     bool              `db:"is_late"`
	CreatedAt  time.Time         `db:"created_at"`
	UpdatedAt  time.Time         `db:"updated_at"`
}

func NewBorrowRequest(req BorrowRequest, book BookEntity) TransactionEntity {
	startDate := time.Now()

	endDate := startDate.AddDate(0, 0, req.Duration)
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 0, 0, 0, endDate.Location())

	return TransactionEntity{
		UserId:    req.UserId,
		BookId:    book.Id,
		Duration:  req.Duration,
		Status:    TransactionStatus_Done,
		StartDate: startDate,
		EndDate:   endDate,
	}
}

func NewReturnRequest(req ReturnRequest, trxEntity TransactionEntity) TransactionEntity {
	rDate := time.Now()

	return TransactionEntity{
		Id:         trxEntity.Id,
		UserId:     req.UserId,
		BookId:     trxEntity.BookId,
		StartDate:  trxEntity.StartDate,
		EndDate:    trxEntity.EndDate,
		ReturnDate: &rDate,
	}
}

func (t *TransactionEntity) SetIsLate() {
	t.IsLate = t.EndDate.After(*t.ReturnDate)
}
