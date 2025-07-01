package database

import (
	"os"

	// Sqlite driver based on CGO

	// "github.com/mattn/go-sqlite3"
	"github.com/glebarez/sqlite"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"

	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details

var Database *gorm.DB = nil

const DATABASE_PATH string = "/opt/vaultrdb/data/store.db"

// build a connection to the local database
func Connect() {

	logging.Default.Info("connecting to local datastore", "src", DATABASE_PATH)

	// creating database file, if not exists
	if _, err := os.Stat(DATABASE_PATH); err != nil {

		logging.Default.Info("couldn't find database -> creating", "destination", DATABASE_PATH)

		if file, err := os.Create(DATABASE_PATH); err != nil {

			logging.Default.Error("couldn't create database file in specified destination", "destination", DATABASE_PATH, "error", err.Error())

			//termination.Shutdown()

		} else {

			file.Close()
		}
	}

	// connecting to the database
	if db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{}); err != nil {

		logging.Default.Error("error connecting to database", "error", err.Error())

		//termination.Shutdown()

	} else {

		Database = db
	}
}

// close database connection if neccessary
func Disconnect() {

	if Database == nil {
		return
	}

	if db, err := Database.DB(); err != nil {

		logging.Default.Error("error receiving database connection", "error", err.Error())
	} else {
		if err := db.Close(); err != nil {

			logging.Default.Error("error closing database connection", "error", err.Error())
		}
	}
}
