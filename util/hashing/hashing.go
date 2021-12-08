package hashing

import (
	"golang.org/x/crypto/bcrypt"
)

func Generate(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(res), err
}

func CompareHash(password, hash_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash_password), []byte(password))
	if err != nil {
		return false
	}
	return true
}
