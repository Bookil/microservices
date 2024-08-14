package domain

type UserID = string

type Auth struct {
	UserID              UserID
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
