package api

import (
	"github.com/gin-gonic/gin"
)

// This function will contain all your endpoint
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Vehicles
		api.GET("/vehicle-makes", handleVehicleMakeList)
	}
}
