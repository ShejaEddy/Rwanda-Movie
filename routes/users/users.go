package users

import (
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/routes"
)

func GetRoutes(db *gorm.DB) (s map[string]routes.SubRoutePackage) {
	s = map[string]routes.SubRoutePackage{
		"/users": {
			Routes: routes.Routes{
				routes.Route{
					Name:    "Create Users",
					Method:  "POST",
					Path:    "/create",
					Handler: nil,
				},
			},
			Middleware: nil,
		},
	}
	return
}
