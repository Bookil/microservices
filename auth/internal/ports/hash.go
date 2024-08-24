package ports

type HashManager interface{
	HashPassword(password string) (hashedPassword string, err error)
	CheckPasswordHash(password, hashedPassword string) bool
}