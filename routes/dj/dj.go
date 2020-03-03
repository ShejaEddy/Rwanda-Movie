package dj

import (
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/controllers/dj"
	"github.com/projects/rwanda-movie/routes"
)

func GetRoutes(db *gorm.DB) (s map[string]routes.SubRoutePackage) {
	s = map[string]routes.SubRoutePackage{
		"/dj": {
			Routes: routes.Routes{
				routes.Route{
					Name:    "Create Dj",
					Method:  "POST",
					Handler: dj.Create(db),
				},
			},
			Middleware: nil,
		},
	}
	return
}
