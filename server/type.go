package server

import "net/http"

type Server struct {
	Addr       string
	Port       int
	HTTPServer *http.Server
}
