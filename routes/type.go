package routes

import "net/http"

type Routes []Route

type Route struct {
	Name       string
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware func(next http.Handler) http.Handler
}

type SubRoutePackage struct {
	Routes     Routes
	Middleware func(next http.Handler) http.Handler
}
