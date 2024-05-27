package objects

import (
	"errors"
	"strings"
	"time"
	"vrdb-storage/server"

	"gorm.io/gorm"
	"vrdb.go/logging"
)

type ServiceConfig struct {
	ID        uint      `json:"id" gorm:"primaryKey,autoIncrement"`
	Key       string    `json:"key" gorm:"<-:create,unique"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create"` // Automatically managed by GORM for creation time
	UpdatedAt time.Time `json:"updated_at"`                  // Automatically managed by GORM for update time
}

// ------------------------------------------------

func (sc *ServiceConfig) BeforeCreate(tx *gorm.DB) (err error) {

	logging.Log.Info("executing *ServiceConfig.BeforeCreate(*gorm.DB)")

	// does the ServiceConfig already exist
	var tmp = ServiceConfig{}
	if result := tx.First(&tmp, "key = ?", sc.Key); result.Error != nil {

		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {

			logging.Log.Info("serviceconfig with given key already exists", "sc", *sc, "tmp", tmp)

			err = result.Error
		}
	}

	return
}

// ------------------------------------------------

// get the value of a specific service configuration
func GetServiceConfig(key string) (string, bool, error) {

	logging.Log.Info("requested key from service_configs", "key", key)

	var sc = ServiceConfig{}
	if result := server.Database.First(&sc, "key = ?", key); result.Error != nil {

		logging.Log.WithValues(
			"key", key,
			"result.RowsAffected", result.RowsAffected).Info("error receiving value from serviceconfig, either request failed or serviceconfig does not exist")

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return "", false, nil
		}

		return "", false, result.Error
	}

	return sc.Value, true, nil
}

// set the value of a specific service configuration
func SetServiceConfig(key, value string) error {

	logging.Log.Info("creating key for service_configs", "key", key, "value", strings.Repeat("*", len(value)))

	var sc = ServiceConfig{
		Key:   key,
		Value: value,
	}

	if result := server.Database.Create(&sc); result.Error != nil {

		logging.Log.WithValues(
			"key", key,
			"value", strings.Repeat("*", len(value)),
			"result.Error", result.Error).Info("error creating key for service_configs")

		return result.Error
	}

	return nil
}
