package services

import (
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/models"
)

func UserExists(db *gorm.DB, user *models.User) (bool, error) {

	users := []models.User{}

	err := db.Where("id = ?", user.ID).Find(&users).Error
	if err != nil {
		return false, err
	}
	if len(users) > 0 {
		return true, nil
	}
	return false, nil
}
