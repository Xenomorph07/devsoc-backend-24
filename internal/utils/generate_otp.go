package utils

import "crypto/rand"

func GenerateOTP(length int) (string, error) {
	code := make([]byte, length)
	if _, err := rand.Read(code); err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		code[i] = uint8(48 + (code[i] % 10))
	}

	return string(code), nil
}
