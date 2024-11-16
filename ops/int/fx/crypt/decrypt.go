package crypt

import "github.com/jnnkrdb/vaultrdb/libs/cryptography"

// decrypt the value with the stored encryption key
func Decrypt(value string) (string, error) {
	return getEncKey(value, cryptography.Decrypt)
}
