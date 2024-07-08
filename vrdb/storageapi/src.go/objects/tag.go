package objects

type Tag struct {
	ID           uint          `json:"-" gorm:"primaryKey,autoIncrement"`
	Tag          string        `json:"tag" gorm:"unique"`
	KeyValueSets []KeyValueSet `json:"-" gorm:"many2many:keyvalueset_tags;"`
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
