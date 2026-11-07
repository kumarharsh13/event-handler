package main

import (
	"event-handler/db"
	event "event-handler/models"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	server.Run(":" + port)
}

func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, event.GetAllEvents())
}

func createEvents(c *gin.Context) {
	var event event.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Event Created!", "event": event})

}
