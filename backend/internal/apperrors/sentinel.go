package apperrors

import "errors"

var (
	ErrInternalServer   = errors.New("internal server error")
	ErrEndpointNotFound = errors.New("endpoint not found")
	ErrEOF              = errors.New("unexpected end of JSON input")
	ErrInvalidToken 		= errors.New("invalid access token")
)