package routes

import (
	"net/http"

	"github.com/Soyaib10/eba-event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func signUp(c *gin.Context) {
	var users models.User

	err := c.ShouldBindJSON(&users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
	}

	err = users.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}