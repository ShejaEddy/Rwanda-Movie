package services

import (
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/models"
	"github.com/satori/go.uuid"
)

func GetDjsPost(db *gorm.DB, userId uuid.UUID) (Post []*models.Post) {

	err := db.Where("user_foreign_key = ?", userId).Find(&Post).Error

	if err != nil {
		panic(err)
		return
	}

	return

}
