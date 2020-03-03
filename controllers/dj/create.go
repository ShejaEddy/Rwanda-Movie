package dj

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

func Create(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
