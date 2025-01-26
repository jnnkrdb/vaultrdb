package vault

import (
	"fmt"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/configstore"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/configstore/buckets"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
	"go.etcd.io/bbolt"
)

// internal func to get the encKey from internal db
func GetConfigByKey(key string) (value string) {

	if err := configstore.DB.View(func(tx *bbolt.Tx) error {

		if value = string(tx.Bucket([]byte(buckets.Vault)).Get([]byte("vrdb-encryption-key"))); value == "" {

			return fmt.Errorf("empty vrdb-encryption-key")
		}

		return nil

	}); err != nil {

		logging.SLog.Error("error receiving value from internal db", "key", key, "error", err.Error())
	}

	return
}
