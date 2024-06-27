package api

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"strings"
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
