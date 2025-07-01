package conf

import (
	"fmt"
	"os"
	"time"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
	"github.com/jnnkrdb/vaultrdb/pkg/storage"
	bolt "github.com/jnnkrdb/vaultrdb/pkg/storage/bbolt"
	"github.com/jnnkrdb/vaultrdb/pkg/storage/cache"
	"go.etcd.io/bbolt"
)

var (
	Vault   storage.Storage
	Authz   storage.Storage
	Configs storage.Storage
)

// load all required storage backends
func LoadStorage() {

	logging.Default.Info("loading storage backends")

	getStorageBackend := func(db string, storageType string) storage.Storage {

		logging.Default.Debug("alloc storage",
			"db", db,
			"storageType", storageType,
		)

		switch storageType {

		case "cache":
			logging.Default.Warn("Storage type CACHE is insufficient for persistency. If you want to keep the stored data, please use another storage type.")
			return cache.NewCacheMap()

		case "bbolt":
			var path = fmt.Sprintf("/opt/vaultrdb/data/%s.db", db)
			if db, err := bolt.NewBBoltDB(
				path,
				&bbolt.Options{
					Timeout:  time.Duration(30 * time.Second),
					ReadOnly: false,
				},
			); err != nil {
				logging.Default.Error("could not open bolt.db",
					"location", path,
				)
				os.Exit(1)
				return nil
			} else {
				return db
			}

		default:
			logging.Default.Error("the configuration for the storage is not valid. you have to decide which implementation to use.")
			os.Exit(1)
			return nil
		}
	}

	if YC.Authz.Enabled {
		Authz = getStorageBackend("authz", YC.Vault.Storage.Type)
	}
	if YC.Configs.Enabled {
		Configs = getStorageBackend("configs", YC.Vault.Storage.Type)
	}
	if YC.Vault.Enabled {
		Vault = getStorageBackend("vault", YC.Vault.Storage.Type)
	}
}
