package response

import (
	"errors"
	"net/http"
)

// Error General
var (
	ErrNotFound = errors.New("data not found")
)
var (
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("email is required")
	ErrPasswordInvalidLength = errors.New("password must have minimun 6 character")
	ErrEmailExist            = errors.New("email already exists")
	ErrPasswordNotMatch      = errors.New("password not match")
)

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(msg, code string, httpcode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpcode,
	}
}
func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral    = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorBadRequest = NewError("bad request", "40000", http.StatusBadRequest)
)

var (
	ErrorEmailRequired         = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)
	ErrorEmailExist            = NewError(ErrEmailExist.Error(), "40901", http.StatusConflict)
	ErrorPasswordNotMatch      = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)
	ErrorNotFound              = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
)

var (
	ErrorMapping = map[string]Error{
		ErrNotFound.Error():              ErrorNotFound,
		ErrEmailRequired.Error():         ErrorEmailRequired,
		ErrEmailInvalid.Error():          ErrorEmailInvalid,
		ErrPasswordRequired.Error():      ErrorPasswordRequired,
		ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
		ErrEmailExist.Error():            ErrorEmailExist,
		ErrPasswordNotMatch.Error():      ErrorPasswordNotMatch,
	}
)
