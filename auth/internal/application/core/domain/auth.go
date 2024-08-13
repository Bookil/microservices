package domain

type UserID = string

type Auth struct {
	UserID              UserID 
	HashedPassword      string
	FailedLoginAttempts int
	AccountLockedUntil  int64
	EmailVerified       bool
	
}

func NewAuthDomain(userID UserID,hashedPassword string)*Auth{
	return &Auth{
		UserID: userID,
		HashedPassword: hashedPassword,
	}
}