package api

type VehicleMake struct {
	ID   int64  `db:"vehicle_make_id" json:"id"`
	Name string `db:"name" json:"name"`
	Url  string `db:"url" json:"url"`
}

type Vehicule struct {
	ID    int64  `db:"vehicle_id" json:"id"`
	Make  int64  `db:"vehicle_make_id" json:"make"`
	Model string `db:"vehicle_model_id" json:"model"`
	Year  int16  `db:"vehicle_year" json:"year"`
}

type VehiculeModel struct {
	ID   int64  `db:"vehicle_model_id" json:"id"`
	Name string `db:"name" json:"name"`
}
