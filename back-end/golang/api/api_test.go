package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

// ─────────────────────────────────────────────────────────────────────────────
// Test helpers
// ─────────────────────────────────────────────────────────────────────────────

func init() {
	gin.SetMode(gin.TestMode)
}

// testRouter returns a Gin engine wired exactly like production (via RegisterRoutes).
func testRouter() *gin.Engine {
	r := gin.New()
	RegisterRoutes(r)
	return r
}

// makeToken generates a valid signed token for use in test requests.
func makeToken(t *testing.T) string {
	t.Helper()
	tok, err := GenerateToken("admin")
	require.NoError(t, err)
	return tok
}

// makeExpiredToken creates a token that is already past its expiry.
func makeExpiredToken(t *testing.T) string {
	t.Helper()
	claims := JWTClaims{
		Username: "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := tok.SignedString(jwtSecret())
	require.NoError(t, err)
	return signed
}

// authHeader returns an Authorization header value for the given token.
func authHeader(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}

// jsonBody serialises v and returns a *bytes.Buffer ready for an http.Request.
func jsonBody(t *testing.T, v any) *bytes.Buffer {
	t.Helper()
	b, err := json.Marshal(v)
	require.NoError(t, err)
	return bytes.NewBuffer(b)
}

// ─────────────────────────────────────────────────────────────────────────────
// POST /api/login
// ─────────────────────────────────────────────────────────────────────────────

// mockUserDB swaps dbGetUserByUserName for the duration of a test.
// It creates a real bcrypt hash of the given password so checkPasswordHash works correctly.
func mockUserDB(t *testing.T, username, password string) func() {
	t.Helper()
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	require.NoError(t, err)

	original := dbGetUserByUsername
	dbGetUserByUsername = func(u string) (*User, error) {
		if u == username {
			return &User{ID: 1, Username: username, PasswordHash: string(hash)}, nil
		}
		return nil, errors.New("user not found")
	}
	return func() { dbGetUserByUsername = original }
}

func TestLogin_ValidCredentials_Returns200WithToken(t *testing.T) {
	defer mockUserDB(t, "admin", "secret")()
	r := testRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/login",
		jsonBody(t, map[string]string{"username": "admin", "password": "secret"}))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp APIResponse[LoginData]
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.True(t, resp.Success)
	assert.NotEmpty(t, resp.Data.Token, "token should not be empty")
}

func TestLogin_WrongPassword_Returns401(t *testing.T) {
	defer mockUserDB(t, "admin", "secret")()
	r := testRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/login",
		jsonBody(t, map[string]string{"username": "admin", "password": "wrong"}))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	var resp APIResponse[any]
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.False(t, resp.Success)
}

func TestLogin_UnknownUser_Returns401(t *testing.T) {
	defer mockUserDB(t, "admin", "secret")()
	r := testRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/login",
		jsonBody(t, map[string]string{"username": "nobody", "password": "secret"}))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestLogin_MissingPassword_Returns400(t *testing.T) {
	r := testRouter() // no DB mock needed — binding fails before the DB is touched

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/login",
		jsonBody(t, map[string]string{"username": "admin"}))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogin_EmptyBody_Returns400(t *testing.T) {
	r := testRouter() // no DB mock needed — binding fails before the DB is touched

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBufferString("{}"))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// ─────────────────────────────────────────────────────────────────────────────
// AuthMiddleware
// ─────────────────────────────────────────────────────────────────────────────

func TestAuthMiddleware_NoHeader_Returns401(t *testing.T) {
	r := testRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/vehicle-makes", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddleware_ValidToken_PassesThrough(t *testing.T) {
	// Swap DB call so the handler succeeds without a real database
	original := dbGetListVehicleMake
	dbGetListVehicleMake = func() ([]*VehicleMake, error) {
		return []*VehicleMake{{ID: 1, Name: "Acura", Url: "acura"}}, nil
	}
	defer func() { dbGetListVehicleMake = original }()

	r := testRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/vehicle-makes", nil)
	req.Header.Set("Authorization", authHeader(makeToken(t)))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthMiddleware_ExpiredToken_Returns401(t *testing.T) {
	r := testRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/vehicle-makes", nil)
	req.Header.Set("Authorization", authHeader(makeExpiredToken(t)))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddleware_TamperedToken_Returns401(t *testing.T) {
	r := testRouter()
	tok := makeToken(t)
	tampered := tok[:len(tok)-4] + "XXXX"

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/vehicle-makes", nil)
	req.Header.Set("Authorization", authHeader(tampered))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddleware_MalformedHeader_Returns401(t *testing.T) {
	r := testRouter()

	cases := []string{
		"token-only",
		"Basic somebase64",
		"Bearer",
	}
	for _, h := range cases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/api/vehicle-makes", nil)
		req.Header.Set("Authorization", h)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code, "header value: %q", h)
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Vehicle handler tests (DB-free via dbFunc variables)
// ─────────────────────────────────────────────────────────────────────────────

func TestHandleVehicleMakeList_Success(t *testing.T) {
	original := dbGetListVehicleMake
	dbGetListVehicleMake = func() ([]*VehicleMake, error) {
		return []*VehicleMake{
			{ID: 1, Name: "Acura", Url: "acura"},
			{ID: 2, Name: "Honda", Url: "honda"},
		}, nil
	}
	defer func() { dbGetListVehicleMake = original }()

	r := testRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/vehicle-makes", nil)
	req.Header.Set("Authorization", authHeader(makeToken(t)))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp APIResponse[VehicleMakeData]
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.True(t, resp.Success)
	assert.Len(t, resp.Data.VehicleMakes, 2)
	assert.Equal(t, "Acura", resp.Data.VehicleMakes[0].Name)
}

func TestHandleVehicleMakeList_DBError_Returns500(t *testing.T) {
	original := dbGetListVehicleMake
	dbGetListVehicleMake = func() ([]*VehicleMake, error) {
		return nil, errors.New("connection refused")
	}
	defer func() { dbGetListVehicleMake = original }()

	r := testRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/vehicle-makes", nil)
	req.Header.Set("Authorization", authHeader(makeToken(t)))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var resp APIResponse[any]
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.False(t, resp.Success)
}

func TestHandleVehicleModelList_Success(t *testing.T) {
	original := dbGetListVehicleModel
	dbGetListVehicleModel = func(makeId string) ([]*VehiculeModel, error) {
		assert.Equal(t, "1", makeId) // verify the route param is forwarded
		return []*VehiculeModel{
			{ID: 10, Name: "ILX"},
			{ID: 11, Name: "MDX"},
		}, nil
	}
	defer func() { dbGetListVehicleModel = original }()

	r := testRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/vehicle-makes/1/vehicle-models", nil)
	req.Header.Set("Authorization", authHeader(makeToken(t)))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp APIResponse[VehicleModelData]
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.True(t, resp.Success)
	assert.Len(t, resp.Data.VehicleModels, 2)
}

func TestHandleVehicleYearsList_Success(t *testing.T) {
	original := dbGetListVehicleYears
	dbGetListVehicleYears = func(makeId string) ([]int16, error) {
		return []int16{2017, 2016, 2015, 2014}, nil
	}
	defer func() { dbGetListVehicleYears = original }()

	r := testRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/vehicle-makes/1/years", nil)
	req.Header.Set("Authorization", authHeader(makeToken(t)))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp APIResponse[VehicleYearData]
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.True(t, resp.Success)
	assert.Equal(t, []int16{2017, 2016, 2015, 2014}, resp.Data.VehicleYears)
}

func TestHandleVehicleMakeCoverageList_Success(t *testing.T) {
	original := dbGetListVehicleMakeCoverage
	dbGetListVehicleMakeCoverage = func(makeId string) (map[string][]int16, error) {
		return map[string][]int16{
			"ILX": {2017, 2016, 2015},
			"MDX": {2017, 2014},
		}, nil
	}
	defer func() { dbGetListVehicleMakeCoverage = original }()

	r := testRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/vehicle-makes/1/coverage", nil)
	req.Header.Set("Authorization", authHeader(makeToken(t)))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp APIResponse[VehicleCoverageData]
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.True(t, resp.Success)
	assert.Contains(t, resp.Data.Coverage, "ILX")
	assert.Contains(t, resp.Data.Coverage, "MDX")
}

func TestHandleCoverageSwitchState_Success(t *testing.T) {
	original := dbSwitchVehicleMakeCoverageState
	dbSwitchVehicleMakeCoverageState = func(makeId, modelId, year string) error {
		assert.Equal(t, "1", makeId)
		assert.Equal(t, "10", modelId)
		assert.Equal(t, "2017", year)
		return nil
	}
	defer func() { dbSwitchVehicleMakeCoverageState = original }()

	r := testRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/vehicle-makes/1/change-coverage/10/2017", nil)
	req.Header.Set("Authorization", authHeader(makeToken(t)))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp APIResponse[PostMessageData]
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.True(t, resp.Success)
	assert.Equal(t, "Coverage state changed successfully", resp.Data.Message)
}

func TestHandleCoverageSwitchState_DBError_Returns500(t *testing.T) {
	original := dbSwitchVehicleMakeCoverageState
	dbSwitchVehicleMakeCoverageState = func(makeId, modelId, year string) error {
		return errors.New("deadlock detected")
	}
	defer func() { dbSwitchVehicleMakeCoverageState = original }()

	r := testRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/vehicle-makes/1/change-coverage/10/2017", nil)
	req.Header.Set("Authorization", authHeader(makeToken(t)))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
