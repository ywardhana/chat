package errormessage

import "errors"

var (
	ErrorFailedAuth = errors.New("Failed Auth")
	ErrNotFound     = errors.New("Not Found")
)
