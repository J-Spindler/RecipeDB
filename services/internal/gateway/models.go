package gateway

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        string         `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type User struct {
	Base
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if len(base.ID) == 0 {
		base.ID = uuid.New().String()
	}
	return
}
