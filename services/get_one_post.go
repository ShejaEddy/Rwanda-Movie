package services

import (
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/models"
	"github.com/satori/go.uuid"
)

func GetOnePost(db *gorm.DB, postId uuid.UUID) (Post []*models.Post) {

	err := db.Preload("User").Where("id = ?", postId).Find(&Post).Error

	if err != nil {
		panic(err)
		return
	}

	return

}
