package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID  `json:"id,omitempty"`
	CreatedAt time.Time  `json:",omitempty"`
	UpdatedAt time.Time  `json:",omitempty"`
	DeletedAt *time.Time `json:",omitempty"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}
