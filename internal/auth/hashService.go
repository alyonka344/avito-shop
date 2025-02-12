package auth

type HashService interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hash, password string) error
}
