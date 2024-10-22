package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password []byte) []byte {
	encryptedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return encryptedPassword
}

func ComparePassword(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	return err == nil
}
