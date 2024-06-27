package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
