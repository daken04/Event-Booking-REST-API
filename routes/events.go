package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/daken04/Event-Booking-REST-API/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"),10,64)
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{
			"message" : "Could not parse event ID",
		})
		return
	}

	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{
			"message" : "Could not get event",
		})
	}
	context.JSON(http.StatusOK,event)
}

func getEvents(context *gin.Context){
	events, err := models.GetAllEvents()

	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError,gin.H{
			"Message" : "Could not parse requested data",
		})
		return
	}

	context.JSON(http.StatusOK,events)
}

func createEvent(context *gin.Context){
	// to access incoming request data
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusBadRequest,gin.H{
			"Message" : "Did not get required fields",
		})
		return
	}

	event.UserID = 1

	err = event.Save()
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError,gin.H{
			"Message" : "Could not create event",
		})
		return
	}

	context.JSON(http.StatusCreated,gin.H{
		"Message" : "Event created successfully",
		"Event" : event,
	})
}

func updateEvent(context *gin.Context){
	eventID, err := strconv.ParseInt(context.Param("id"),10,64)
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{
			"message" : "Cannot parse id parameter",
		})
		return
	}

	_, err = models.GetEvent(eventID)
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{
			"message" : "Event id does not exist",
		})
		return
	}

	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{
			"message" : "Did not get required fields",
		})
	}

	updateEvent.ID = eventID
	err = updateEvent.Update()

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{
			"message" : "Not able to update",
		})
	}

	context.JSON(http.StatusCreated,gin.H{
		"Message" : "Event created successfully",
		"Event" : updateEvent,
	})
}