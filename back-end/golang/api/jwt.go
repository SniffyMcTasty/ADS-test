package api

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// jwtSecret is loaded from the JWT_SECRET env var (set it in your .env file).
// Falls back to a dev placeholder so the server still starts locally.
func jwtSecret() []byte {
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		return []byte(secret)
	}
	return []byte("dev-secret-change-me")
}

// JWTClaims defines the payload stored inside each token.
type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken creates a signed HS256 JWT valid for 24 hours.
func GenerateToken(username string) (string, error) {
	claims := JWTClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "adsdata-api",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret())
}

// AuthMiddleware validates the Bearer token on every protected route.
// On success it forwards the username via c.Set("username", ...).
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, APIResponse[any]{
				Success: false,
				Error:   &APIErrorResponse{Code: 401, Message: "Authorization header is required"},
			})
			return
		}

		// Expect "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, APIResponse[any]{
				Success: false,
				Error:   &APIErrorResponse{Code: 401, Message: "Authorization header format must be: Bearer <token>"},
			})
			return
		}

		claims := &JWTClaims{}
		token, err := jwt.ParseWithClaims(parts[1], claims, func(t *jwt.Token) (interface{}, error) {
			// Guard against algorithm-confusion attacks
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret(), nil
		})

		if err != nil || token == nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, APIResponse[any]{
				Success: false,
				Error:   &APIErrorResponse{Code: 401, Message: "Invalid or expired token"},
			})
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
