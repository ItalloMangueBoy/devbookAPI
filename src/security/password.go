package security

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword: Gen password hash
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// ValidPassword: Compares a hashed password with its possible plaintext equivalent.
func ValidPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
