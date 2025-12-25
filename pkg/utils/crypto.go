package utils

import "golang.org/x/crypto/bcrypt"

// CryptoHash generates a bcrypt hash of the password combined with salt.
// It returns the hashed password and any error encountered during hashing.
func CryptoHash(raw, salt string) (string, error) {
	combined := raw + salt
	hash, err := bcrypt.GenerateFromPassword([]byte(combined), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CryptoHashCompare compares a raw password with salt against a bcrypt hash.
// It returns true if the password matches the hash, false otherwise.
func CryptoHashCompare(raw, salt, hash string) bool {
	combined := raw + salt
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(combined))
	return err == nil
}
