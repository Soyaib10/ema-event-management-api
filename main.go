package main

import (
	"net/http"
	"strconv"

	"github.com/Soyaib10/eba-event-booking-api/db"
	"github.com/Soyaib10/eba-event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()

	r.GET("/events", getEvents)
	r.GET("/events/:id", getEvent)
	r.POST("/events", createEvent)

	r.Run(":8080") 
}

func getEvent(c *gin.Context) {
	// Get id from path
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fectch event. Try again later."})
		return
	}
	c.JSON(http.StatusOK, event)
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fectch events. Try again later."})
		return
	}
	c.JSON(http.StatusOK, events) // auto return with JSON formate
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse JSON data"})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create an evnet. Tey again later."})
	// }
	c.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}