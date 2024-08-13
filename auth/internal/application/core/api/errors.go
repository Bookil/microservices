package api

import "errors"

var (
	ErrHashingPassword  = errors.New("error hashing password")
	ErrCreateAuthStore  = errors.New("error create auth store")
	ErrCreateEmailToken = errors.New("error create email token")
)
