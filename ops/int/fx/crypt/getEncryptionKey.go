package crypt

import (
	"fmt"

	"github.com/jnnkrdb/vaultrdb/int/configstore"
	"github.com/jnnkrdb/vaultrdb/int/configstore/buckets"
	"go.etcd.io/bbolt"
)

// internal func to get the encKey from internal db
func getEncKey(value string, f func(string, string) (string, error)) (string, error) {

	var encKey string

	if err := configstore.DB.View(func(tx *bbolt.Tx) error {

		encKey = string(tx.Bucket([]byte(buckets.Vault)).Get([]byte("vrdb-encryption-key")))

		if encKey = string(tx.Bucket([]byte(buckets.Vault)).Get([]byte("vrdb-encryption-key"))); encKey == "" {

			return fmt.Errorf("error empty vrdb-encryption-key")
		}

		return nil

	}); err != nil {

		return "", err
	}

	return f(encKey, value)
}
