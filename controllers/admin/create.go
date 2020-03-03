package admin

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/helpers"
	"github.com/projects/rwanda-movie/models"
	"github.com/projects/rwanda-movie/services"
	"net/http"
)

func Create(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var creds models.User

		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			services.SendError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if ok, errors := helpers.ValidateInputs(creds); !ok {
			helpers.ValidationResponse(errors, w)
			return
		}

		var users []models.User

		users, err = services.FindByEmail(db, &creds)
		if err != nil {
			panic(err)
			return
		}

		if len(users) > 0 {
			services.SendError(w, http.StatusBadRequest, "User Already Exists")
		} else {

			result := services.CreateAdmin(&creds, db)

			if !result.Status {
				services.SendError(w, http.StatusBadRequest, result.Message)
			} else {
				services.SendSuccess(w, http.StatusOK, result.Message, result.Data)
			}

		}

	}
}
