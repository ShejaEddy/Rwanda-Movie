package posts

import (
	"encoding/json"
	"fmt"
	"github.com/alexedwards/scs"
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/helpers"
	"github.com/projects/rwanda-movie/models"
	"github.com/projects/rwanda-movie/services"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"net/http"
)

func createPost(db *gorm.DB, uploads *models.Post, w http.ResponseWriter) {
	err := db.Create(&uploads).Error
	if err != nil {
		services.SendError(w, http.StatusInternalServerError, err.Error())
	} else {
		services.SendSuccess(w, http.StatusOK, "Post Created Successfully", uploads)
	}
}

func Create(db *gorm.DB, session *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var uploads models.Post

		log.Println(session.Get(r.Context(), "role"))
		err := json.NewDecoder(r.Body).Decode(&uploads)
		if err != nil {
			services.SendError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if ok, errors := helpers.ValidateInputs(uploads); !ok {
			helpers.ValidationResponse(errors, w)
			return
		}
		uploads.UserID = uuid.FromStringOrNil(session.GetString(r.Context(), "userID"))
		if uploads.Mode == "normal" {
			createPost(db, &uploads, w)
		} else {
			r.ParseMultipartForm(10 << 20)

			file, _, err := r.FormFile("image")
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				return
			}
			defer file.Close()

			var uploads models.Post

			tempFile, err := ioutil.TempFile("uploads", "image-*.png")
			if err != nil {
				fmt.Println(err)
			}

			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			tempFile.Write(fileBytes)

			if tempFile.Name() != "" {
				uploads.Image = tempFile.Name()
			}
			defer tempFile.Close()

			createPost(db, &uploads, w)
		}

	}
}
