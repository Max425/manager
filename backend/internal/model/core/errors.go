package core

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrInternal = errors.New("internal error")
	ErrRequired = errors.New("required parameter is omitted")
)
