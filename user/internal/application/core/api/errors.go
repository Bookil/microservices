package api

import "errors"

var (
	ErrRegisterFailed           = errors.New("registering failed please try again")
	ErrEmailRegistered          = errors.New("someone registered with this email already")
	ErrUserNotFindWithThisEmail = errors.New("there is no user with this email")
	ErrLoggingFailed            = errors.New("logging failed try again please")
	ErrAccessDenied             = errors.New("access denied")
	ErrChangingPasswordFailed   = errors.New("changing password failed please try again")
	ErrResetPasswordFailed      = errors.New("reset password failed please try again")
	ErrUpdateFailed             = errors.New("updating user failed please try again")
	ErrDeleteAccountFailed      = errors.New("deleting account failed please try again")
)
