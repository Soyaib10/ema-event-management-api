package routes

import "github.com/gin-gonic/gin"

// This file is responsible for rotues
func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", getEvents)
	r.GET("/events/:id", getEvent)
	r.POST("/events", createEvent)
	r.PUT("/events/:id", updateEvent)
	r.DELETE("/events/:id", deleteEvent)
}