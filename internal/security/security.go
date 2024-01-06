package security

import "golang.org/x/crypto/bcrypt"

func CheckPassword(passwordHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}
