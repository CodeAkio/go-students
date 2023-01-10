package main

import (
	"github.com/CodeAkio/go-students/database"
	"github.com/CodeAkio/go-students/routes"
)

func main() {
	database.ConnectDb()
	routes.HandleRequests()
}
