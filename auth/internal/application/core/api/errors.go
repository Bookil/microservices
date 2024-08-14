package api

import "errors"

var (
	ErrHashingPassword              = errors.New("error hashing password")
	ErrCreateAuthStore              = errors.New("error create auth store")
	ErrCreateEmailToken             = errors.New("error create email token")
	ErrAccessDenied                 = errors.New("error access denied")
	ErrNotFound                     = errors.New("error not found")
	ErrInvalidEmailPassword         = errors.New("invalid email or password")
	ErrInvalidPassword              = errors.New("invalid password")
	ErrDelete                       = errors.New("error delete auth")
	ErrChangePassword               = errors.New("error change password")
	ErrVerifyEmail                  = errors.New("error verify email")
	ErrDestroyToken                 = errors.New("error destroy token")
	ErrLockAccount                  = errors.New("error lock account")
	ErrAccountLocked                = errors.New("error account locked try again after 2 min")
	ErrIncrementFailedLoginAttempts = errors.New("error increment failed login attempts")
	ErrEmailNotVerified             = errors.New("error not verified")
	ErrClearFailedLoginAttempts     = errors.New("error clear failed login attempts")
	ErrGenerateToken                = errors.New("error generate token")
)
