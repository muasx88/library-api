package transaction

type BorrowRequest struct {
	UserId   int
	BookIsbn string `json:"book_isbn" validate:"required"`
	Duration int    `json:"duration" validate:"required"`
}

type ReturnRequest struct {
	UserId   int
	BookIsbn string `json:"book_isbn" validate:"required"`
}
