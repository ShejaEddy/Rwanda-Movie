package admin

import (
	"encoding/json"
	"gopkg.in/go-playground/validator.v10"
	"net/http"
	"reflect"
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

		reflected := reflect.ValueOf(dataSet)
		for _, err := range err.(validator.ValidationErrors) {

			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string

			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors[name] = append(errors[name], "The "+name+" is required")
				break
			case "email":
				errors[name] = append(errors[name], "The "+name+" should be a valid email")
				break
			case "min":
				errors[name] = append(errors[name], "The "+name+" should be a greater than 6 characters")
				break
			case "eqfield":
				errors[name] = append(errors[name], "The "+name+" should be equal to the "+err.Param())
				break
			default:
				errors[name] = append(errors[name], "The "+name+" is invalid")
				break
			}
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
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(message)

}
