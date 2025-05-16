package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run()
}

func getEvents(context *gin.Context){
	context.JSON(http.StatusOK,gin.H{
		"Message":"Hello",
	})
}