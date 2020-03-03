package services

import (
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/models"
)

func FindByEmail(db *gorm.DB, user *models.User) (users []models.User, error error) {

	err := db.Where("email = ?", user.Email).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return
}
