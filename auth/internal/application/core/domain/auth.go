package domain

import "time"

type UserID = string

type Role = string

const(
	UserRole Role = "user"
	AdminRole Role = "admin"
)

type Auth struct {
	UserID              UserID `gorm:"unique"`
	Role Role
	HashedPassword      string
	FailedLoginAttempts int
	AccountLockedUntil  int64
	IsEmailVerified     bool
}

func NewAuth(userID UserID,role Role,hashedPassword string) *Auth {
	return &Auth{
		Role :role,
		UserID:         userID,
		HashedPassword: hashedPassword,
	}
}

type AccessTokenClaims struct {
	UserID   string
	Role Role
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
