package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GetHashedPassword returns hashed password
func GetHashedPassword(fromPwd string) string {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(fromPwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPwd)
}

// CompareHashedPasswords compares two passwords equals or not
func CompareHashedPasswords(fromPwd, checkPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(fromPwd), []byte(checkPwd))
	return err != nil
}
