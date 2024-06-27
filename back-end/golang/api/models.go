package api

type VehicleMake struct {
	ID   int64  `db:"vehicle_make_id" json:"id"`
	Name string `db:"name" json:"name"`
	Url  string `db:"url" json:"url"`
}
