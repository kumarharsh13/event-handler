package main

import (
	"event-handler/db"
	event "event-handler/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8000") //localhost:8000
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
