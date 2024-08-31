package domain

import "time"

type UserID = string

type Auth struct {
	UserID              UserID `gorm:"unique"`
	HashedPassword      string
	FailedLoginAttempts int
	AccountLockedUntil  int64
	IsEmailVerified     bool
}

func NewAuth(userID UserID, hashedPassword string) *Auth {
	return &Auth{
		UserID:         userID,
		HashedPassword: hashedPassword,
	}
}

type AccessTokenClaims struct {
	UserID   string
	CreateAt time.Time
}

type RefreshTokenClaims struct {
	IPAddress  string
	UserAgent  string
	LoggedInAt time.Duration
}

type ResetPasswordTokenClaims struct {
	UserID   string
	CreateAt time.Time
}
