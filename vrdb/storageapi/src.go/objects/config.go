package objects

import "time"

const (
	TYPE_Config string = "config"
	TYPE_Secret string = "secret"
)

type KeyValuePair struct {
	ID          uint   `json:"id"`
	Key         string `json:"key" gorm:"<-:create"`
	Value       string `json:"value"`
	Type        string `json:"type" gorm:"<-:create"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"created_at" gorm:"<-:create"` // Automatically managed by GORM for creation time
	UpdatedAt time.Time `json:"updated_at"`                  // Automatically managed by GORM for update time
}
