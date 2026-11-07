package main

import (
	"event-handler/db"
	"event-handler/routes"
	"fmt"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	routes.RegisterRoutes(server)

	server.Run(":" + port)
}
