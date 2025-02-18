package autherrors

import "errors"

var (
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
	ErrEmailExists = errors.New("email already have an account")
)