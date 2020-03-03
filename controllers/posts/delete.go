package posts

import (
	"github.com/alexedwards/scs"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/services"
	"github.com/satori/go.uuid"
	"net/http"
)

func DeleteOne(db *gorm.DB, session *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		postId := uuid.FromStringOrNil(params["postId"])

		ok := services.DeletePost(db, postId)
		if ok {
			services.SendSuccess(w, http.StatusOK, "Post Successfully Deleted.", nil)
		} else {
			services.SendError(w, http.StatusBadRequest, "Post Failed To Be Deleted.")
		}

	}
}
