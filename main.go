package main

import (
	DBFactory "github.com/projects/rwanda-movie/db"
	serverFactory "github.com/projects/rwanda-movie/server"
	"os"
	"strconv"
)

var port int
var dbHost string
var dbPort string
var dbUser string
var dbPassword string
var dbDatabase string

func init() {

	rawport := os.Getenv("PORT")

	if len(rawport) > 0 {
		var err error
		port, err = strconv.Atoi(rawport)
		if err != nil {
			panic(err)
			return
		}
	} else {

		port = 6500
		dbHost = "localhost"
		dbPort = "3306"
		dbUser = "root"
		dbPassword = ""
		dbDatabase = "rwanda-movie"

	}
}

func main() {
	db, err := DBFactory.Connect(dbHost, dbPort, dbUser, dbPassword, dbDatabase)
	defer db.Close()

	if err != nil {
		panic(err)
	}
	server := serverFactory.NewServer(port, db)
	server.Start()
}
