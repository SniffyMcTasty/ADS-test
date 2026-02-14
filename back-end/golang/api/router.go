package api

import (
	"github.com/gin-gonic/gin"
)

// This function will contain all your endpoint
func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")

	// Public routes (no token required)
	api.POST("/login", handleLogin)

	// Protected routes (valid JWT required)
	protected := api.Group("/", AuthMiddleware())
	{
		// Vehicles
		protected.GET("/vehicle-makes", handleVehicleMakeList)
		protected.GET("/vehicle-makes/:makeId/vehicle-models", handleVehicleModelList)
		protected.GET("/vehicle-makes/:makeId/years", handleVehicleYearsList)
		protected.GET("/vehicle-makes/:makeId/coverage", handleVehicleMakeCoverageList)
		protected.POST("/vehicle-makes/:makeId/change-coverage/:modelId/:year", handleVehicleMakeCoverageSwitchState)
	}
}
