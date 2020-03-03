package models

import (
	"github.com/satori/go.uuid"
)

type Post struct {
	Base
	Mode        string    `gorm:"default:'normal'" validate:"required" json:"mode"`
	Title       string    `validate:"" json:"name,omitempty"`
	Genre       string    `validate:"" json:"genre,omitempty"`
	Image       string    `validate:"" json:"image,omitempty"`
	Content     string    `validate:"" json:"content,omitempty"`
	Description string    `validate:"" json:"description,omitempty"`
	Type        string    `validate:"" json:"type,omitempty"`
	Duration    string    `json:"duration"`
	UserID      uuid.UUID `gorm:"index;column:user_foreign_key;not null" json:"userId,omitempty"`
	User        User      `json:"poster_data,omitempty"`
}
