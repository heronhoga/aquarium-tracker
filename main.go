package main

import (
	"go-restapi-gin-aquarium/models"
	"go-restapi-gin-aquarium/routes"
)

func main() {
	models.ConnectDatabase()
	routes.Route()
	
}