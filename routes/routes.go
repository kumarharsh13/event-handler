package routes

import (
	"event-handler/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Event Routes
	server.GET("/events/:id", getEvent)
	server.GET("/events", getEvents)

	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)

	authenticate.POST("/events", middlewares.Authenticate, createEvent)
	authenticate.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	authenticate.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)

	//THIS another way
	// server.POST("/events", middlewares.Authenticate, createEvent)
	// server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	// server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)

	// Auth Routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
