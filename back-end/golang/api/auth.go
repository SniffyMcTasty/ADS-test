package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Login
// @Description Authenticate and receive a JWT token
// @Tags Auth
// @Accept  json
// @Produce json
// @Param   credentials body LoginRequest true "Username and password"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} APIErrorResponse
// @Failure 401 {object} APIErrorResponse
// @Failure 500 {object} APIErrorResponse
// @Router /api/login [post]
func handleLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, APIResponse[any]{
			Success: false,
			Error:   &APIErrorResponse{Code: 400, Message: "username and password are required"},
		})
		return
	}

	// Credential check
	user, err := dbGetUserByUsername(req.Username)
	if err != nil || !checkPasswordHash(req.Password, user.PasswordHash) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, APIResponse[any]{
			Success: false,
			Error:   &APIErrorResponse{Code: 401, Message: "Invalid credentials"},
		})
		return
	}

	token, err := GenerateToken(req.Username)
	if err != nil {
		log.Printf("Cannot generate token. err=%v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, APIResponse[any]{
			Success: false,
			Error:   &APIErrorResponse{Code: 500, Message: "Could not generate token"},
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse[LoginData]{
		Success: true,
		Data:    LoginData{Token: token},
	})
}

func checkPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
