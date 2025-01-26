package authstore

import (
	"time"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
	"github.com/jnnkrdb/vaultrdb/pkg/storedb"
	"go.etcd.io/bbolt"
)

const (
	Path = "/opt/vaultrdb/data/authstore.db"
)

var (
	DB *storedb.LocalDB = &storedb.LocalDB{}
)

func InitDB() {

	logging.SLog.Info("opening authstore", "path", Path)

	if err := DB.OpenDB(Path, &bbolt.Options{
		Timeout: 10 * time.Second,
	}); err != nil {

		logging.SLog.Error("error opening database", "err", err)
	}
}
