package middleware

import (
	"github.com/alexedwards/scs"
	"github.com/casbin/casbin"
	"github.com/projects/rwanda-movie/services"
	"net/http"
)

func Authorizer(e *casbin.Enforcer, session *scs.SessionManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			role := session.GetString(r.Context(), "role")

			if role == "" {
				role = "anonymous"
			}

			res, err := e.Enforce(role, r.URL.Path, r.Method)
			if err != nil {
				services.SendError(w, http.StatusInternalServerError, "Error Enforcing Authorization")
				return
			}
			if res {
				next.ServeHTTP(w, r)
			} else {
				services.SendError(w, http.StatusForbidden, "Forbidden Route")
				return
			}
		}

		return http.HandlerFunc(fn)
	}
}
