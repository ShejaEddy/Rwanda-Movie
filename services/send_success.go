package services

import (
	"encoding/json"
	"github.com/projects/rwanda-movie/dtos"
	"net/http"
)

func SendSuccess(w http.ResponseWriter, status int, message string, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	result := dtos.Response{Status: true, Message: message, Code: status, Data: data}

	packet, err := json.Marshal(result)

	if err != nil {
		panic(err)
		return
	}

	w.Write(packet)
}
