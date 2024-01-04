package transaction

import "time"

type UserHistoryTransactionResponse struct {
	Id         int               `json:"id"`
	UserId     int               `json:"user_id"`
	BookId     int               `json:"book_id"`
	BookTitle  string            `json:"book_title"`
	BookIsbn   string            `json:"book_isbn"`
	Duration   int               `json:"duration"`
	StartDate  time.Time         `json:"start_date"`
	EndDate    time.Time         `json:"end_date"`
	ReturnDate time.Time         `json:"return_date"`
	Status     TransactionStatus `json:"status"`
	IsLate     bool              `json:"is_late"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

func NewUserHistoryTransactionResponse(transactions []TransactionEntity) (res []UserHistoryTransactionResponse) {
	for _, t := range transactions {
		res = append(res, UserHistoryTransactionResponse{
			Id:         t.Id,
			UserId:     t.UserId,
			BookId:     t.BookId,
			BookTitle:  t.BookTitle,
			BookIsbn:   t.BookIsbn,
			Duration:   t.Duration,
			StartDate:  t.StartDate,
			EndDate:    t.EndDate,
			ReturnDate: *t.ReturnDate,
			Status:     t.Status,
			IsLate:     t.IsLate,
			CreatedAt:  t.CreatedAt,
			UpdatedAt:  t.UpdatedAt,
		})
	}

	return
}
