package helpers

import (
	"encoding/json"
	"github.com/projects/rwanda-movie/langs"
	"gopkg.in/go-playground/validator.v10"
	"net/http"
	"strings"
)

func ValidateInputs(dataSet interface{}) (bool, map[string][]string) {

	var validate *validator.Validate
	validate = validator.New()

	err := validate.Struct(dataSet)

	if err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		errors := make(map[string][]string)

		for _, err := range err.(validator.ValidationErrors) {
			var name string
			name = strings.ToLower(err.StructField())
			langs.GenerateValidationMessage(name, err.Tag(), errors)
		}

		return false, errors
	}
	return true, nil
}

func ValidationResponse(fields map[string][]string, w http.ResponseWriter) {

	response := make(map[string]interface{})
	response["status"] = "error"
	response["message"] = "validation error"
	response["errors"] = fields
	message, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occured internally"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(message)

}
