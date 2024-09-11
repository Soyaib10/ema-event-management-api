package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/events", getEvents)
	r.Run(":8080") 
}

func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}