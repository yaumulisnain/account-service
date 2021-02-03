package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 8

func HashPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcryptCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
