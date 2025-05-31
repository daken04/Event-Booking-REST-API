package main

import (
	"github.com/daken04/Event-Booking-REST-API/routes"
	"github.com/daken04/Event-Booking-REST-API/db"
	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}

