package api

type VehicleMake struct {
	ID   int64  `db:"vehicle_make_id" json:"id"`
	Name string `db:"name" json:"name"`
	Url  string `db:"url" json:"url"`
}

type Vehicule struct {
	ID      int64  `db:"vehicle_id" json:"id"`
	Make    int64  `db:"vehicle_make_id" json:"make"`
	Model   string `db:"vehicle_model_id" json:"model"`
	Year    int16  `db:"vehicle_year" json:"year"`
	State   int16  `db:"state" json:"state"`
	Updated string `db:"updated" json:"updated"`
}

type VehiculeModel struct {
	ID   int64  `db:"vehicle_model_id" json:"id"`
	Name string `db:"name" json:"name"`
}

type User struct {
	ID           int64  `db:"user_id" json:"id"`
	Username     string `db:"username" json:"username"`
	PasswordHash string `db:"password_hash" json:"-"`
}

// Data structures for API responses
type VehicleMakeData struct {
	VehicleMakes []*VehicleMake `json:"vehicle_makes"`
}

type VehicleModelData struct {
	VehicleModels []*VehiculeModel `json:"vehicle_models"`
}

type VehicleYearData struct {
	VehicleYears []int16 `json:"vehicle_years"`
}

type VehicleCoverageData struct {
	Coverage map[string][]int16 `json:"coverage"`
}

type PostMessageData struct {
	Message string `json:"message"`
}

// Response structures for API endpoints
type APIErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type APIResponse[T any] struct {
	Success bool              `json:"success"`
	Data    T                 `json:"data,omitempty"`
	Error   *APIErrorResponse `json:"error,omitempty"`
}

// Response structures for Swagger documentation
type VehicleMakeResponse struct {
	Success bool            `json:"success"`
	Data    VehicleMakeData `json:"data"`
}

type VehicleModelResponse struct {
	Success bool             `json:"success"`
	Data    VehicleModelData `json:"data"`
}

type VehicleYearResponse struct {
	Success bool            `json:"success"`
	Data    VehicleYearData `json:"data"`
}

type VehicleCoverageResponse struct {
	Success bool                `json:"success"`
	Data    VehicleCoverageData `json:"data"`
}

type PostMessageResponse struct {
	Success bool            `json:"success"`
	Data    PostMessageData `json:"data"`
}

// Auth types
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginData struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Success bool      `json:"success"`
	Data    LoginData `json:"data"`
}
