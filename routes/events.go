package routes

import (
	"net/http"
	"strconv"

	"github.com/Soyaib10/eba-event-booking-api/models"
	"github.com/Soyaib10/eba-event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

// This file is responsible for making application handler.
func deleteEvent(c *gin.Context) {
	// Get id from path
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	// Get all info of a single id
	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event successfully deleted"})
}

func updateEvent(c *gin.Context) {
	// Get id from path
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	// Get all info of a single id
	_, err = models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update successfull"})
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
	// Authorization
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	// normal part
	var event models.Event
	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse JSON data"})
		return
	}

	event.UserID = 1
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}
