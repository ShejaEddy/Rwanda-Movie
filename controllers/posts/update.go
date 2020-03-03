package posts

import (
	"encoding/json"
	"github.com/alexedwards/scs"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/models"
	"github.com/projects/rwanda-movie/services"
	"github.com/satori/go.uuid"
	"net/http"
)

func UpdateOne(db *gorm.DB, session *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		postId := uuid.FromStringOrNil(params["postId"])

		var post models.Post

		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			services.SendError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		post.ID = postId
		ok := services.UpdatePost(db, post)
		if ok {
			services.SendSuccess(w, http.StatusOK, "Post Successfully Updated.", nil)
		} else {
			services.SendError(w, http.StatusBadRequest, "Post Failed To Be Updated.")
		}

	}
}
