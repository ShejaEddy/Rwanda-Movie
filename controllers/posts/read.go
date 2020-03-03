package posts

import (
	"github.com/alexedwards/scs"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/services"
	"github.com/satori/go.uuid"
	"net/http"
)

func GetAll(db *gorm.DB, session *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		posts := services.GetAllPosts(db)
		services.SendSuccess(w, http.StatusOK, "List Of All Posts", posts)
	}
}

func GetOne(db *gorm.DB, session *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		postId := uuid.FromStringOrNil(params["postId"])

		posts := services.GetOnePost(db, postId)
		if len(posts) > 0 {
			services.SendSuccess(w, http.StatusOK, "Post ID Detail", posts[0])
		} else {
			services.SendSuccess(w, http.StatusOK, "Post ID Detail", posts)
		}
	}
}

func GetDjsPost(db *gorm.DB, session *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		userId := uuid.FromStringOrNil(params["userId"])
		posts := services.GetDjsPost(db, userId)
		services.SendSuccess(w, http.StatusOK, "List Of All Posts From DJ", posts)
	}
}
