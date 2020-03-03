package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/models"
)

func Save(user *models.User, db *gorm.DB) (r RepositoryResult) {
	err := db.Save(user).Error
	if err != nil {
		r.Error = err
		return
	}
	r.Result = user
	return
}
