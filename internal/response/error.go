package response

import (
	"errors"
	"net/http"
)

var (
	ErrInternalServerError = errors.New("internal Server Error")
	ErrConflict            = errors.New("your Item already exist")
	ErrBadParamInput       = errors.New("given Param is not valid")
)

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrUserAlreadyExists     = errors.New("email already exists")
	ErrUserPasswordNotMatch  = errors.New("email or password is incorrect")
	ErrUserNotFoundSendEmail = errors.New("silahkan cek email anda, jika email terdaftar maka akan menerima link reset password")
)

var (
	ErrPasswordRequired        = errors.New("password is required")
	ErrNewPasswordRequired     = errors.New("new password is required")
	ErrConfirmPasswordRequired = errors.New("confirm password is required")
	ErrConfirmPasswordNotMatch = errors.New("password or confirm password not match")
)

var (
	ErrBookNotFound      = errors.New("book not found")
	ErrBookAlreadyExists = errors.New("book already exists")
	ErrBookStock         = errors.New("book stock invalid")
)

var (
	ErrTransactionNotFound        = errors.New("transaction not found")
	ErrForbiddenTransaction       = errors.New("forbidden transaction. user must return the book")
	ErrForbiddenReturnTransaction = errors.New("forbidden transaction. user never borrow the book")
	ErrWrongReturnTransaction     = errors.New("forbidden transaction. invalid book to return")
)

var ErrorMap = map[error]int{
	ErrInternalServerError:        http.StatusInternalServerError,
	ErrConflict:                   http.StatusConflict,
	ErrUserNotFound:               http.StatusNotFound,
	ErrUserPasswordNotMatch:       http.StatusBadRequest,
	ErrUserAlreadyExists:          http.StatusBadRequest,
	ErrBookNotFound:               http.StatusNotFound,
	ErrBookAlreadyExists:          http.StatusBadRequest,
	ErrTransactionNotFound:        http.StatusNotFound,
	ErrForbiddenTransaction:       http.StatusBadRequest,
	ErrForbiddenReturnTransaction: http.StatusBadRequest,
	ErrWrongReturnTransaction:     http.StatusBadRequest,
}

type IError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value,omitempty"`
	Message string `json:"message"`
}
