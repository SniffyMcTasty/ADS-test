package main

import (
	"adsdata.ca/backendapi/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create new router (gin.Engine)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Initialize the routes
	api.RegisterRoutes(r)

	// Start serving the application
	err = r.Run(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))

	if err != nil {
		log.Fatal("Error Starting Web Server")
	}
}