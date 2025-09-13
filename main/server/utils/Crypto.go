package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (*string, error) {
	const cost = 8

	passwordBytes := []byte(password)

	hashed, err := bcrypt.GenerateFromPassword(passwordBytes, cost)
	if err != nil {
		return nil, err
	}

	hashedPassword := string(hashed)
	return &hashedPassword, nil
}

func ComparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
