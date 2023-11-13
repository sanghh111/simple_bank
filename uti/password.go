package uti

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword	returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	} else {
		return string(hashPassword), err
	}
}

// CheckPassword return the bool correct passord
func CheckPassword(password string, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
