package routes

import (
	"event-handler/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userID := c.GetInt64("userID")
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	eventFound, err := models.GetEventById(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the event"})
	}

	err = eventFound.Register(userID)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register in the event"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Thank You for registering", "event": eventFound})

}

func cancleRegistration(c *gin.Context) {
	userID := c.GetInt64("userID")
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	eventFound, err := models.GetEventById(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the event"})
	}

	err = eventFound.CancelRegistration(eventFound.ID, userID)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to cancel from event"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Sorry to see you go...", "event": eventFound})
}
