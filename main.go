package main

import (
	"net/http"

	"github.com/Soyaib10/eba-event-booking-api/db"
	"github.com/Soyaib10/eba-event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()

	r.GET("/events", getEvents)
	r.POST("/events", createEvent)
	r.Run(":8080") 
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
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
	c.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}