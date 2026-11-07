package routes

import (
	event "event-handler/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	foundEvent, err := event.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, foundEvent)
}

func getEvents(c *gin.Context) {
	events, err := event.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch data..!!!. Try again later."})
	}
	c.JSON(http.StatusOK, events)
}

func createEvents(c *gin.Context) {
	var event event.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	event.UserID = 1

	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't insert data..!!!. Try again later."})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event Created!", "event": event})

}
