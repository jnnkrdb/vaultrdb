package funcs

import (
	"fmt"
	"vrdb-storage/server/configs"

	"go.etcd.io/bbolt"
	"vrdb.go/cryptography"
)

// encrypt the value with the stored encryption key
func Encrypt(value string) (string, error) {
	return getEncKey(value, cryptography.Encrypt)
}

// decrypt the value with the stored encryption key
func Decrypt(value string) (string, error) {
	return getEncKey(value, cryptography.Decrypt)
}

// internal func to get the encKey from internal db
func getEncKey(value string, f func(string, string) (string, error)) (string, error) {

	var encKey string

	if err := configs.DB.View(func(tx *bbolt.Tx) error {

		encKey = string(tx.Bucket([]byte(configs.BucketVault)).Get([]byte("vrdb-encryption-key")))

		if encKey = string(tx.Bucket([]byte(configs.BucketVault)).Get([]byte("vrdb-encryption-key"))); encKey == "" {

			return fmt.Errorf("error empty vrdb-encryption-key")
		}

		return nil

	}); err != nil {

		return "", err
	}

	return f(encKey, value)
}
