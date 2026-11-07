package routes

import (
	user "event-handler/models"
	"event-handler/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var u user.User

	err := c.ShouldBindJSON(&u)
	fmt.Print(err)
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

func login(c *gin.Context) {
	var login user.LoginRequest

	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	err = login.ValidateCredentials()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(login.Email, login.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't authenticate USER"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login Successfull", "token": token})
}
