package controllers

import (
	"net/http"
)

func Welcome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome To Rwanda-Movie API."))
	}
}
