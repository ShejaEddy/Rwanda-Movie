package server

import (
	"github.com/alexedwards/scs"
	"github.com/casbin/casbin"
	"github.com/gorilla/handlers"
	"github.com/jinzhu/gorm"
	routerFactory "github.com/projects/rwanda-movie/handlers"
	"github.com/projects/rwanda-movie/middleware"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func (s *Server) Start() {
	log.Println("server has started on port", s.Port)
	log.Fatal(s.HTTPServer.ListenAndServe())
}

func NewServer(port int, db *gorm.DB) *Server {
	var s Server

	s.Port = port
	s.Addr = ":" + strconv.Itoa(port)

	var session *scs.SessionManager

	session = scs.New()
	session.IdleTimeout = 30 * time.Minute

	router := routerFactory.NewRouter(db, session)

	authEnforcer, err := casbin.NewEnforcer("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	routes := session.LoadAndSave(middleware.Authorizer(authEnforcer, session)(router.Router))

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "X-Access-Token", "Cashe-Control"}),
		handlers.ExposedHeaders([]string{}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(routes))

	s.HTTPServer = &http.Server{
		Addr:           s.Addr,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &s
}
