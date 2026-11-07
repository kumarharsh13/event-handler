package routes

import (
	user "event-handler/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var u user.User

	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	err = u.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to SignUp. Try again later."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Sign-Up Successful"})
}
