package controllers

import (
	"encoding/json"
	"github.com/alexedwards/scs"
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/jwt"
	"github.com/projects/rwanda-movie/models"
	"github.com/projects/rwanda-movie/services"
	"github.com/projects/rwanda-movie/tools/password"
	"net/http"
)

func Login(db *gorm.DB, session *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var creds models.User

		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			services.SendError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		users, err := services.FindByEmail(db, &creds)

		if err != nil {
			panic(err)
			return
		}

		if len(users) > 0 {
			user := users[0]
			if !password.Compare(user.Password, creds.Password) {
				services.SendError(w, http.StatusUnauthorized, "Invalid Credentials")
				return
			} else {
				token, err := jwt.CreateToken(user.ID)
				if err != nil {
					services.SendError(w, http.StatusInternalServerError, err.Error())
					return
				}
				if err := session.RenewToken(r.Context()); err != nil {
					services.SendError(w, http.StatusInternalServerError, "Can Not Renew Token")
					return
				}
				session.Put(r.Context(), "userID", user.ID.String())
				session.Put(r.Context(), "role", user.Role)
				type success struct {
					Token string
				}
				services.SendSuccess(w, http.StatusOK, "User Successfully Authenticated", success{Token: token})
			}

		} else {
			services.SendError(w, http.StatusUnauthorized, "Invalid Credentials")
			return
		}

	}
}

func Logout(db *gorm.DB, session *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session.Remove(r.Context(), "role")
		session.Remove(r.Context(), "userID")
		services.SendSuccess(w, http.StatusOK, "user Successfully LoggedOut", nil)
	}
}
