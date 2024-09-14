package routes

import (
	"github.com/Soyaib10/eba-event-booking-api/middlewares"
	"github.com/gin-gonic/gin"
)

// This file is responsible for rotues
func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", getEvents)
	r.GET("/events/:id", getEvent)

	authenticate := r.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.POST("/events", createEvent)
	authenticate.PUT("/events/:id", updateEvent)
	authenticate.DELETE("/events/:id", deleteEvent)
	authenticate.POST("/events/:id/register", registerForEvents)
	authenticate.DELETE("/events/:id/register", cancleRegistration)

	r.POST("/signup", signUp)
	r.POST("/login", login)
}