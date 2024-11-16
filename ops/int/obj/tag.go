package obj

import (
	"github.com/jnnkrdb/vaultrdb/libs/logging"
	"gorm.io/gorm"
)

type Tag struct {
	ID    uint   `json:"-" gorm:"primaryKey,autoIncrement"`
	KvsID uint   `json:"-"`
	Tag   string `json:"tag"`
}

// ------------------------------------------------

//func (tag *Tag) BeforeCreate(tx *gorm.DB) (err error) {
//
//	logging.Log.Info("executing *Tag.BeforeCreate(*gorm.DB)")
//
//	// does the tag already exist
//	var tmp = Tag{}
//	if result := server.Database.First(&tmp, "tag = ?", tag.Tag); result.Error == nil {
//
//		if result.RowsAffected != 0 {
//
//			logging.Log.Info("tag with given key already exists", "tag", *tag, "tmp", tmp)
//
//			err = fmt.Errorf("tag already exists")
//		}
//
//	} else {
//
//		if errors.Is(gorm.ErrRecordNotFound, result.Error) {
//			return nil
//		}
//
//		err = result.Error
//	}
//
//	return
//}

func (tag *Tag) AfterCreate(tx *gorm.DB) (err error) {

	logging.Log.Info("executing *Tag.AfterCreate(*gorm.DB)")

	if err := tx.Delete(&Tag{}, "kvs_id IS NULL").Error; err != nil {

		logging.Log.Info("error removing unused tags")
	}

	return
}
