package utils

import (
	_rand "crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateSalt(size int) (string, error) {
	salt := make([]byte, size)
	n, err := _rand.Read(salt)
	if err != nil {
		return "", err
	}
	if n != size {
		return "", fmt.Errorf("failed to generate salt: expected %d bytes, got %d", size, n)
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}
