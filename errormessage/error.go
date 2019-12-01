package errormessage

import "errors"

var (
	ErrorFailedAuth = errors.New("Failed Auth")
	ErrorUnexpected = errors.New("Unexpected Error")
	ErrNotFound     = errors.New("Not Found")
)
