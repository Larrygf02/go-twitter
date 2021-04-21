package bd

import (
	"github.com/go-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.User, bool) {
	user, found, _ := ExistUser(email)
	if !found {
		return user, false
	}
	passwordCrypt := []byte(password)
	passwordBD := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordCrypt)
	if err != nil {
		return user, false
	}
	return user, true
}
