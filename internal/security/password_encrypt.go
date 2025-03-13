package security

import ("golang.org/x/crypto/bcrypt")
func EncryptPassword(password string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}

func VerifyPassword (password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

