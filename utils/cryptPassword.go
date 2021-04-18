package utils

import "golang.org/x/crypto/bcrypt"

func CryptPassword(password string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost) // mayor costo mayor demora
	return string(bytes), err
}
