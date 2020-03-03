package services

import (
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/dtos"
	"github.com/projects/rwanda-movie/models"
	"github.com/projects/rwanda-movie/repositories"
	"github.com/projects/rwanda-movie/tools/password"
	"net/http"
)

func CreateAdmin(user *models.User, db *gorm.DB) (r dtos.Response) {

	var w http.ResponseWriter
	var err error
	user.Password, err = password.Encrypt(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Role = "admin"

	result := repositories.Save(user, db)

	if result.Error != nil {
		r.Status = false
		r.Message = result.Error.Error()
		return
	}

	r.Status = true
	r.Message = "User Successfully Created"
	r.Data = result.Result.(*models.User)
	return
}
