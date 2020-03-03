package routes

import (
	"github.com/alexedwards/scs"
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/controllers"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		},
	)
}

func GetRoutes(db *gorm.DB, session *scs.SessionManager) Routes {
	return Routes{
		{
			Name:    "Home",
			Method:  "GET",
			Path:    "/",
			Handler: controllers.Welcome(),
		},
		{
			Name:    "Users Authentication",
			Method:  "POST",
			Path:    "/auth",
			Handler: controllers.Login(db, session),
		},
		{
			Name:    "Users Logout",
			Method:  "POST",
			Path:    "/logout",
			Handler: controllers.Logout(db, session),
		},
	}
}
