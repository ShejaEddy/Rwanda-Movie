package services

import (
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/models"
)

func UpdatePost(db *gorm.DB, post models.Post) bool {

	err := db.Save(&post).Error

	if err != nil {
		panic(err)
		return false
	}

	return true

}
