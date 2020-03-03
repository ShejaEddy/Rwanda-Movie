package services

import (
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/models"
)

func GetAllPosts(db *gorm.DB) (Post []*models.Post) {

	err := db.Preload("User").Find(&Post).Error

	if err != nil {
		panic(err)
		return
	}

	return

}
