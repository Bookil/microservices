package mysql_adapter

import "errors"

var ErrBookNotFound = errors.New("book not found")
var ErrGenreNotFound = errors.New("genre not found")
var ErrAuthorNotFound = errors.New("author not found")