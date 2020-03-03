package admin

import (
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/controllers/admin"
	"github.com/projects/rwanda-movie/routes"
)

func GetRoutes(db *gorm.DB) (s map[string]routes.SubRoutePackage) {
	s = map[string]routes.SubRoutePackage{
		"/admin": {
			Routes: routes.Routes{
				routes.Route{
					Name:    "Create Admin",
					Method:  "POST",
					Handler: admin.Create(db),
				},
				routes.Route{
					Name:    "Create Admin",
					Method:  "GET",
					Handler: admin.Create(db),
				},
			},
			Middleware: nil,
		},
	}
	return
}
