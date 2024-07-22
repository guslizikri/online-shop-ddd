package response

import "errors"

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
