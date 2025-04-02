package middlewares

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"api-order/src/shared/responses"
)

var jwtKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, responses.Response{Success: false,
				Message: "acceso denegado para el recurso solicitado",
				Error:   "token no proporcionado o invalido el token proporcionado"})
			c.Abort()
			return
		}

		tokenString := authHeader[len("Bearer "):]

		claims := &CustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, responses.Response{
				Success: false,
				Message: "acceso denegado para el recurso solicitado",
				Error:   "invalido el token proporcionado"})
			c.Abort()
			return
		}
		c.Set("datUser", claims)
		c.Next()
	}
}
