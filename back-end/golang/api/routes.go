package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleVehicleMakeList(c *gin.Context) {

	vehicleMakes, err := GetListVehicleMake()

	if err != nil {
		log.Printf("Cannot load vehicle make. err=%v", err)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "Cannot load vehicle make"},
		)

		return
	}

	c.JSON(http.StatusOK, gin.H{"vehicle_makes": vehicleMakes})
}

func handleVehicleModelList(c *gin.Context) {

	vehicleModels, err := GetListVehicleModel(c.Param("makeId"))

	if err != nil {
		log.Printf("Cannot load vehicle models. err=%v", err)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "Cannot load vehicle models"},
		)

		return
	}

	c.JSON(http.StatusOK, gin.H{"vehicle_models": vehicleModels})
}

func handleVehicleYearsList(c *gin.Context) {

	vehicleYears, err := GetListVehicleYears()

	if err != nil {
		log.Printf("Cannot load vehicle years. err=%v", err)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "Cannot load vehicle years"},
		)

		return
	}

	c.JSON(http.StatusOK, gin.H{"vehicle_years": vehicleYears})
}

func handleVehicleMakeCoverageList(c *gin.Context) {

	coverage, err := GetListVehicleMakeCoverage(c.Param("makeId"))

	if err != nil {
		log.Printf("Cannot load coverages. err=%v", err)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "Cannot load coverages"},
		)

		return
	}

	c.JSON(http.StatusOK, gin.H{"coverage": coverage})
}
