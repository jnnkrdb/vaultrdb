package objects

import (
	"fmt"
	"time"
	"vrdb-storage/server"

	"gorm.io/gorm"
	"vrdb.go/logging"
)

//const (
//	TYPE_Config string = "config"
//	TYPE_Secret string = "secret"
//)

type KeyValueSet struct {
	ID    uint   `json:"id" gorm:"primaryKey,autoIncrement"`
	Key   string `json:"key" gorm:"<-:create,unique"`
	Value string `json:"value"`
	//Type        string `json:"type" gorm:"<-:create"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"created_at" gorm:"<-:create"` // Automatically managed by GORM for creation time
	UpdatedAt time.Time `json:"updated_at"`                  // Automatically managed by GORM for update time
}

// ------------------------------------------------

func (kvs *KeyValueSet) BeforeCreate(tx *gorm.DB) (err error) {

	logging.Log.Info("executing *KeyValueSet.BeforeCreate(*gorm.DB)")

	// does the keyvalueset already exist
	var tmp = KeyValueSet{}
	if result := server.Database.First(&tmp, "key = ?", kvs.Key); result.Error == nil {

		if result.RowsAffected != 0 {

			logging.Log.Info("keyvalueset with given key already exists", "kvs", *kvs, "tmp", tmp)

			err = fmt.Errorf("keyvalueset already exists")
		}
	}

	return
}
