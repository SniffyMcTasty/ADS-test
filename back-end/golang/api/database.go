package api

import (
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Example how to connect to the database mysql located in docker
func GetConn() (*sqlx.DB, error) {
	parts := []string{os.Getenv("DB_USER"), ":", os.Getenv("DB_PASSWORD"), "@tcp(", os.Getenv("DB_HOST"), ":", os.Getenv("DB_PORT"), ")/", os.Getenv("DB_SCHEMA")}
	db, err := sqlx.Connect("mysql", strings.Join(parts, ""))

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Example how to get the list of vehicle make with the library sqlx
func GetListVehicleMake() ([]*VehicleMake, error) {
	var listVehicleMake []*VehicleMake

	db, err := GetConn()

	if err != nil {
		return listVehicleMake, err
	}

	defer db.Close()

	err = db.Select(&listVehicleMake, "SELECT * FROM vehicle_make")

	if err != nil {
		return listVehicleMake, err
	}

	return listVehicleMake, nil
}

func GetListVehicleModel(makeId string) ([]*VehiculeModel, error) {
	var listVehicleModel []*VehiculeModel

	db, err := GetConn()

	if err != nil {
		return listVehicleModel, err
	}

	defer db.Close()

	err = db.Select(&listVehicleModel, "SELECT model.* FROM (SELECT DISTINCT vehicle_model_id FROM vehicle WHERE vehicle_make_id = ?) v LEFT JOIN vehicle_model model ON model.vehicle_model_id = v.vehicle_model_id", makeId)

	if err != nil {
		return listVehicleModel, err
	}

	return listVehicleModel, nil
}

func GetListVehicleYears(makeId string) ([]int16, error) {
	var listVehicleYear []int16

	db, err := GetConn()

	if err != nil {
		return listVehicleYear, err
	}

	defer db.Close()

	err = db.Select(&listVehicleYear, "SELECT DISTINCT vehicle_year FROM vehicle WHERE vehicle_make_id = ? ORDER BY vehicle_year DESC", makeId)

	if err != nil {
		return listVehicleYear, err
	}

	return listVehicleYear, nil
}

func GetListVehicleMakeCoverage(makeId string) (map[string][]int16, error) {
	var listVehicleMakeCoverage map[string][]int16 = make(map[string][]int16)

	db, err := GetConn()

	if err != nil {
		return listVehicleMakeCoverage, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT DISTINCT name, vehicle_year FROM vehicle LEFT JOIN vehicle_model ON vehicle.vehicle_model_id = vehicle_model.vehicle_model_id WHERE vehicle.vehicle_make_id = ? and vehicle.state = 1 ORDER BY vehicle_year DESC", makeId)

	if err != nil {
		return listVehicleMakeCoverage, err
	}

	defer rows.Close()

	for rows.Next() {
		var modelName string
		var vehicleYear int16
		err := rows.Scan(&modelName, &vehicleYear)
		if err != nil {
			return listVehicleMakeCoverage, err
		}
		listVehicleMakeCoverage[modelName] = append(listVehicleMakeCoverage[modelName], vehicleYear)
	}
	return listVehicleMakeCoverage, nil
}

func SwitchVehicleMakeCoverageState(makeId string, modelId string, year string) error {
	db, err := GetConn()

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE vehicle SET state = 1 - state WHERE vehicle_make_id = ? AND vehicle_model_id = ? AND vehicle_year = ?", makeId, modelId, year)

	if err != nil {
		return err
	}

	return nil
}
