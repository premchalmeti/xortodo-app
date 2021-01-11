package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// GetHashedPassword returns hashed password
func GetHashedPassword(fromPwd string) string {
	hashedPwd, err := bcrypt.GenerateFromPassword(
		[]byte(fromPwd), bcrypt.DefaultCost,
	)
	if err != nil {
		// panic(err)
		return ""
	}
	return string(hashedPwd)
}

// CompareHashedPasswords compares two passwords equals or not
func CompareHashedPasswords(hashedPwd, checkPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(checkPwd))
	fmt.Println(err)
	return err == nil
}
