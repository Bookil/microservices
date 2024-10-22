package mysql_adapter

import "errors"

var (
	ErrBookNotFound   = errors.New("book not found")
	ErrGenreNotFound  = errors.New("genre not found")
	ErrAuthorNotFound = errors.New("author not found")
)
