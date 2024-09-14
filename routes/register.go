package routes

import (
	"net/http"
	"strconv"

	"github.com/Soyaib10/eba-event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvents(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register((userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user."})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successfull"})
}

func cancleRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err !=  nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id"})
	}

	var event models.Event
	event.ID = eventId
	err = event.CancleRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancle registration."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cancled!"})
}