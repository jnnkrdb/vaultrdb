package server

import (
	"os"

	// Sqlite driver based on CGO
	"vrdb.go/logging"

	// "github.com/mattn/go-sqlite3"
	"github.com/glebarez/sqlite"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details

var Database *gorm.DB

const (
	_DATABASE_PATH string = "/opt/vaultrdb/data/store.db"
)

func ConnectToDatabase() {

	logging.Log.Info("connecting to local datastore", "src", _DATABASE_PATH)

	// creating database file, if not exists
	if _, err := os.Stat(_DATABASE_PATH); err != nil {

		logging.Log.Info("couldn't find database -> creating", "destination", _DATABASE_PATH)

		if file, err := os.Create(_DATABASE_PATH); err != nil {

			logging.Log.Error(err, "couldn't create database file in specified destination", "destination", _DATABASE_PATH)

			os.Exit(1)

		} else {

			file.Close()
		}
	}

	// connecting to the database
	if db, err := gorm.Open(sqlite.Open(_DATABASE_PATH), &gorm.Config{}); err != nil {

		logging.Log.Error(err, "error connecting to database")

		os.Exit(1)

	} else {

		Database = db
	}
}
