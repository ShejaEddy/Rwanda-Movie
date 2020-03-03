package services

import (
	"encoding/json"
	"github.com/projects/rwanda-movie/dtos"
	"net/http"
)

func SendError(w http.ResponseWriter, status int, message string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errors := dtos.Response{Status: false, Message: message, Code: status}

	packet, err := json.Marshal(errors)

	if err != nil {
		panic(err)
		return
	}

	w.Write(packet)
}
