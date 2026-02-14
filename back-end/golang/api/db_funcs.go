package api

// These variables point to the real database functions by default.
// In tests, reassign them to return controlled data without a live DB:
//
//	dbGetListVehicleMake = func() ([]*VehicleMake, error) {
//	    return []*VehicleMake{{ID: 1, Name: "Acura"}}, nil
//	}
var (
	dbGetListVehicleMake             = GetListVehicleMake
	dbGetListVehicleModel            = GetListVehicleModel
	dbGetListVehicleYears            = GetListVehicleYears
	dbGetListVehicleMakeCoverage     = GetListVehicleMakeCoverage
	dbSwitchVehicleMakeCoverageState = SwitchVehicleMakeCoverageState
	dbGetUserByUsername              = GetUserByUsername
)
