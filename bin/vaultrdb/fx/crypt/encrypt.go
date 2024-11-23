package crypt

import "github.com/jnnkrdb/vaultrdb/pkg/cryptography"

// encrypt the value with the stored encryption key
func Encrypt(value string) (string, error) {
	return getEncKey(value, cryptography.Encrypt)
}
