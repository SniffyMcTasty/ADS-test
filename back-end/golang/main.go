// @title           ADS Vehicle Coverage API
// @version         1.0
// @description     API for managing vehicle make, model, year and coverage data.
// @host            localhost:8080
// @schemes         http
// @basePath        /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"log"
	"os"

	"adsdata.ca/backendapi/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
