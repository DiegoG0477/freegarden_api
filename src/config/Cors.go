package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ConfigurationCors() gin.HandlerFunc {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	frontendURL := os.Getenv("FRONTEND_URL")

	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		ExposeHeaders:    []string{"Content-Length"},
	}

	// Permite todos los orígenes con credenciales
	if frontendURL == "*" {
		config.AllowOriginFunc = func(origin string) bool {
			return true
		}
		// Importante: Cuando usas AllowOriginFunc, el header Access-Control-Allow-Origin
		// se establecerá al valor del origen que haga la petición (no se puede usar *)
	} else {
		// Para múltiples URLs específicas
		config.AllowOrigins = strings.Split(frontendURL, ",")
	}

	return cors.New(config)
}
