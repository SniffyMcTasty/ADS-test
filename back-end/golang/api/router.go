package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// This function will contain all your endpoint
func RegisterRoutes(r *gin.Engine) {
	r.Use(corsMiddleware())

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
