package autherrors

import "errors"

var (
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)