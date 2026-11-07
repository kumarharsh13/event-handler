package routes

import (
	event "event-handler/models"
	"event-handler/utils"
	"net/http"
	"strconv"
	"strings"

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

func createEvent(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization header"})
		return
	}

	token := parts[1]

	userID, err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
		return
	}

	var e event.Event
	err = c.ShouldBindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	e.UserID = userID

	err = e.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't insert data. Try again later."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": e})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	_, err = event.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	var updateEvent event.Event

	err = c.ShouldBindJSON(&updateEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	err = updateEvent.Update(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't update data..!!!. Try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event Updated!", "event": updateEvent})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	_, err = event.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	err = event.DeleteEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
