package configstore

import (
	"time"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
	"github.com/jnnkrdb/vaultrdb/pkg/storedb"
	"go.etcd.io/bbolt"
)

const (
	Path = "/opt/vaultrdb/data/configstore.db"
)

var (
	DB *storedb.LocalDB = &storedb.LocalDB{}
)

func InitDB() {

	logging.Default.Info("opening configstore", "path", Path)

	if err := DB.OpenDB(Path, &bbolt.Options{
		Timeout: 10 * time.Second,
	}); err != nil {

		logging.Default.Error("error opening database", "err", err)
	}
}
