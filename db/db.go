package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/projects/rwanda-movie/models"
)

func Connect(
	host string,
	port string,
	user string,
	password string,
	database string,
) (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		return
	}

	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Role{})
	db.LogMode(true)

	return
}
