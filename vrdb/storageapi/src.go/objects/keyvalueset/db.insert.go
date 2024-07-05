package keyvalueset

import (
	"vrdb-storage/objects/tag"
	"vrdb-storage/server"

	"vrdb.go/logging"
)

// insert a new keyvalueset into the database
func InsertIntoDB(newObj NewKeyValueSet) (KeyValueSet, error) {

	// create object in database without the tags, then append the tags later
	var keyvalueset = KeyValueSet{
		Key:         newObj.Key,
		Value:       newObj.Value,
		Tags:        []tag.Tag{},
		Description: newObj.Description,
	}

	if result := server.Database.Create(&keyvalueset); result.Error != nil {

		logging.Log.Info("error creating keyvalueset in database", "newObj", newObj, "result.Error", result.Error)

		return KeyValueSet{}, result.Error
	}

	// get the object from the database
	keyvalueset = KeyValueSet{}

	if result := server.Database.Preload("Tags").First(&keyvalueset, "key = ?", newObj.Key); result.Error != nil {

		logging.Log.Info("error finding kvs", "newObj", newObj, "result.Error", result.Error)

		return KeyValueSet{}, result.Error
	}

	// append the tags to the object, if any
	if len(newObj.Tags) > 0 {

		for i := range newObj.Tags {

			if err := server.Database.Model(&keyvalueset).Association("Tags").Append(&newObj.Tags[i]); err != nil {

				logging.Log.Info("error appending tag to keyvalueset", "kvs", keyvalueset, "tag", newObj.Tags[i], "err", err)

				return KeyValueSet{}, err
			}
		}
	}

	return keyvalueset, nil
}
