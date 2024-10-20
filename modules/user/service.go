package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// using bcrypt = password encryption
func isEncryptedPassword(hashedPassword, password string) (error) {
	match := checkPassword(hashedPassword, password)

	if !match {
		return errors.New("wrong password provided")
	}

	return nil
}

func checkPassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(password),
	)
    return err == nil
}

func hashPassword(password string) (string, error){
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)

    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}