package obj

import (
	"os"
	"reflect"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/database"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// migrate predefined structs into the database
func Migrate() {

	var (
		// contains the names of the objects, which should be migrated
		// into the database
		listOfObjectNames = []string{}

		// contains the items, which schould be migrated into the database
		listOfObjects = []interface{}{
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

	logging.Default.Info("migrating objects into database", "object-list", listOfObjectNames)

	if err := database.Database.AutoMigrate(listOfObjects...); err != nil {

		logging.Default.Error("couldn't migrate objects into database", "error", err.Error())

		os.Exit(1)
	}
}
