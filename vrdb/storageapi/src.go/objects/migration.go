package objects

import (
	"os"
	"reflect"
	"vrdb-storage/server"

	"vrdb.go/logging"
)

// migrate predefined structs into the database
func Migrate() {

	var (
		// contains the names of the objects, which should be migrated
		// into the database
		listOfObjectNames = []string{}

		// contains the items, which schould be migrated into the database
		listOfObjects = []interface{}{
			&Tag{},
			&KeyValueSet{},
		}
	)

	// parse through the list of migrate-objects and get their names, for the log
	for _, t := range listOfObjects {
		if o := reflect.TypeOf(t); o.Kind() == reflect.Ptr {
			listOfObjectNames = append(listOfObjectNames, o.Elem().Name())
		} else {
			listOfObjectNames = append(listOfObjectNames, o.Name())
		}
	}

	logging.Log.Info("migrating objects into database", "object-list", listOfObjectNames)

	if err := server.Database.AutoMigrate(listOfObjects...); err != nil {

		logging.Log.Error(err, "couldn't migrate objects into database")

		os.Exit(1)
	}
}
