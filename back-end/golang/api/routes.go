package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get the list of vehicle make
// @Description Get the list of vehicle make
// @Tags Vehicles
// @Success 200 {object} VehicleMakeResponse
// @Failure 500 {object} APIErrorResponse
// @Router /api/vehicle-makes [get]
func handleVehicleMakeList(c *gin.Context) {

	vehicleMakes, err := dbGetListVehicleMake()

	if err != nil {
		log.Printf("Cannot load vehicle make. err=%v", err)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			APIResponse[any]{
				Success: false,
				Error: &APIErrorResponse{
					Code:    500,
					Message: "Cannot load vehicle make",
				},
			},
		)

		return
	}

	response := APIResponse[VehicleMakeData]{
		Success: true,
		Data:    VehicleMakeData{VehicleMakes: vehicleMakes},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get a list of vehicle models
// @Description Get the list of vehicle models by make id
// @Tags Models
// @Param makeId path int true "Vehicle Make ID"
// @Success 200 {object} VehicleModelResponse
// @Failure 500 {object} APIErrorResponse
// @Router /api/vehicle-makes/:makeId/vehicle-models [get]
func handleVehicleModelList(c *gin.Context) {

	vehicleModels, err := dbGetListVehicleModel(c.Param("makeId"))

	if err != nil {
		log.Printf("Cannot load vehicle models. err=%v", err)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			APIResponse[any]{
				Success: false,
				Error: &APIErrorResponse{
					Code:    500,
					Message: "Cannot load vehicle models",
				},
			},
		)

		return
	}

	response := APIResponse[VehicleModelData]{
		Success: true,
		Data:    VehicleModelData{VehicleModels: vehicleModels},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get a list of vehicle years
// @Description Get the list of vehicle years for every model by make id
// @Tags Years
// @Param makeId path int true "Vehicle Make ID"
// @Success 200 {object} VehicleYearResponse
// @Failure 500 {object} APIErrorResponse
// @Router /api/vehicle-makes/:makeId/years [get]
func handleVehicleYearsList(c *gin.Context) {

	vehicleYears, err := dbGetListVehicleYears(c.Param("makeId"))

	if err != nil {
		log.Printf("Cannot load vehicle years. err=%v", err)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			APIResponse[any]{
				Success: false,
				Error: &APIErrorResponse{
					Code:    500,
					Message: "Cannot load vehicle years",
				},
			},
		)

		return
	}

	response := APIResponse[VehicleYearData]{
		Success: true,
		Data:    VehicleYearData{VehicleYears: vehicleYears},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get a list of vehicle coverage
// @Description Get the list of vehicle coverage for every model and year by make id
// @Tags Coverage
// @Param makeId path int true "Vehicle Make ID"
// @Success 200 {object} VehicleCoverageResponse
// @Failure 500 {object} APIErrorResponse
// @Router /api/vehicle-makes/:makeId/coverage [get]
func handleVehicleMakeCoverageList(c *gin.Context) {

	coverage, err := dbGetListVehicleMakeCoverage(c.Param("makeId"))

	if err != nil {
		log.Printf("Cannot load coverages. err=%v", err)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			APIResponse[any]{
				Success: false,
				Error: &APIErrorResponse{
					Code:    500,
					Message: "Cannot load coverages",
				},
			},
		)

		return
	}

	response := APIResponse[VehicleCoverageData]{
		Success: true,
		Data:    VehicleCoverageData{Coverage: coverage},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Switch the coverage state of a vehicle
// @Description Change the coverage state of a vehicle by make id, model id and year
// @Tags Coverage
// @Param makeId path int true "Vehicle Make ID"
// @Param modelId path int true "Vehicle Model ID"
// @Param year path int true "Vehicle Year"
// @Success 200 {object} PostMessageResponse
// @Failure 500 {object} APIErrorResponse
// @Router /api/vehicle-makes/:makeId/change-coverage/:modelId/:year [post]
func handleVehicleMakeCoverageSwitchState(c *gin.Context) {

	err := dbSwitchVehicleMakeCoverageState(c.Param("makeId"), c.Param("modelId"), c.Param("year"))

	if err != nil {
		log.Printf("Cannot change coverage state. err=%v", err)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			APIResponse[any]{
				Success: false,
				Error: &APIErrorResponse{
					Code:    500,
					Message: "Cannot change coverage state",
				},
			},
		)

		return
	}

	c.JSON(http.StatusOK, APIResponse[PostMessageData]{
		Success: true,
		Data:    PostMessageData{Message: "Coverage state changed successfully"},
	})
}
