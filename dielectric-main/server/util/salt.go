package util

import (
	"golang.org/x/crypto/bcrypt"
)

// Compare a salt hash.
func Compare(digest []byte, password *string) bool {
	hex := []byte(*password)
	if err := bcrypt.CompareHashAndPassword(digest, hex); err == nil {
		return true
	}
	return false
}
