package bd

import (
	"github.com/Neil-uli/tweeto/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.User, bool) {
	us, found, _ := CheckIfExistUser(email)
	if found == false {
		return us, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(us.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return us, false
	}
	return us, true
}
