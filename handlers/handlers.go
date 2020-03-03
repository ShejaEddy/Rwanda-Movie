package handlers

import (
	"github.com/alexedwards/scs"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/routes"
	"github.com/projects/rwanda-movie/routes/admin"
	"github.com/projects/rwanda-movie/routes/dj"
	"github.com/projects/rwanda-movie/routes/posts"
	"github.com/projects/rwanda-movie/routes/users"
)

func (r *RouteHandler) NewSubRoutes(
	path string,
	subroutes routes.Routes,
	Middleware mux.MiddlewareFunc,
) (subrouter *mux.Router) {
	subrouter = r.Router.PathPrefix(path).Subrouter()
	for _, subroute := range subroutes {
		subrouter.
			Methods(subroute.Method).
			Path(subroute.Path).
			Name(subroute.Name).
			Handler(subroute.Handler)

	}
	return
}

func (r *RouteHandler) AddRoutesWithSubRoutes(pack map[string]routes.SubRoutePackage) {
	for path, route := range pack {
		r.NewSubRoutes(path, route.Routes, route.Middleware)
	}
}

func NewRouter(db *gorm.DB, session *scs.SessionManager) *RouteHandler {
	var r RouteHandler

	r.Router = mux.NewRouter().StrictSlash(true)

	Routes := routes.GetRoutes(db, session)

	for _, route := range Routes {
		r.Router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.Handler)
	}

	r.AddRoutesWithSubRoutes(admin.GetRoutes(db))
	r.AddRoutesWithSubRoutes(dj.GetRoutes(db))
	r.AddRoutesWithSubRoutes(users.GetRoutes(db))
	r.AddRoutesWithSubRoutes(posts.GetRoutes(db, session))
	return &r
}
