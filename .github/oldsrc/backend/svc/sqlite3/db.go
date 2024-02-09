package sqlite3

import (
	"database/sql"
	"os"

	"github.com/go-logr/logr"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

const DB string = "/data/vault.db"

var _DBConn *sql.DB

func ConnectDB(_log logr.Logger) {

	var dbIsNew bool = false

	if _, err := os.Stat(DB); err != nil {
		_log.Info("db does not exist and will be created", "source", DB)

		if file, err := os.Create(DB); err != nil {

			_log.Error(err, "error creating db", "source", DB)
			os.Exit(1)
		} else {
			dbIsNew = true
			file.Close()
		}
	}

	_log.Info("connecting to db", "source", DB)

	if c, err := sql.Open("sqlite3", DB); err != nil {
		_log.Error(err, "error connecting to db", "source", DB)
		os.Exit(1)
	} else {
		_DBConn = c
	}

	if dbIsNew {
		if _, err := _DBConn.Exec("CREATE TABLE IF NOT EXISTS vault (uid text NOT NULL PRIMARY KEY, data text NOT NULL);"); err != nil {
			_log.Error(err, "error creating reuqired table", "source", DB, "table", "vault")
			os.Exit(1)
		}
	}
}

// disconnect from postgres
func Disconnect() error {
	return _DBConn.Close()
}
