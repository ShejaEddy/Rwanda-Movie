package posts

import (
	"github.com/alexedwards/scs"
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/controllers/posts"
	"github.com/projects/rwanda-movie/routes"
)

func GetRoutes(db *gorm.DB, session *scs.SessionManager) (s map[string]routes.SubRoutePackage) {
	s = map[string]routes.SubRoutePackage{
		"/posts": {
			Routes: routes.Routes{
				routes.Route{
					Name:    "Create Post",
					Method:  "POST",
					Handler: posts.Create(db, session),
				},
				routes.Route{
					Name:    "Get All Posts",
					Method:  "GET",
					Handler: posts.GetAll(db, session),
				},
				routes.Route{
					Name:    "Get One Post",
					Method:  "GET",
					Path:    "/{postId}",
					Handler: posts.GetOne(db, session),
				},
				routes.Route{
					Name:    "Get Dj's Posts",
					Method:  "GET",
					Path:    "/dj/{userId}",
					Handler: posts.GetDjsPost(db, session),
				},
				routes.Route{
					Name:    "Delete Post",
					Method:  "DELETE",
					Path:    "/{postId}",
					Handler: posts.DeleteOne(db, session),
				},
				routes.Route{
					Name:    "Delete Post",
					Method:  "PUT",
					Path:    "/{postId}",
					Handler: posts.UpdateOne(db, session),
				},
			},
			Middleware: nil,
		},
	}
	return
}
