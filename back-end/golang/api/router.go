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
		api.GET("/vehicle-makes/:makeId/vehicle-models", handleVehicleModelList)
		api.GET("/vehicle-makes/:makeId/years", handleVehicleYearsList)
		api.GET("/vehicle-makes/:makeId/coverage", handleVehicleMakeCoverageList)
		api.POST("/vehicle-makes/:makeId/change-coverage/:modelId/:year", handleVehicleMakeCoverageSwitchState)
	}
}
