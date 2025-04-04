package config

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigurationCors() gin.HandlerFunc {
	config := cors.Config{
		// Permite todos los orígenes
		AllowAllOrigins: true,

		// Permite todos los métodos HTTP
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS",
		},

		// Permite todos los headers comunes y personalizados
		AllowHeaders: []string{
			"Origin", "Content-Length", "Content-Type", "Accept",
			"Authorization", "X-Requested-With", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers", "Access-Control-Allow-Methods",
			"Access-Control-Allow-Credentials", "Accept-Encoding",
			"Accept-Language", "Cache-Control", "Connection", "Cookie",
			"Host", "Pragma", "Referer", "User-Agent", "*",
		},

		// Expone todos los headers
		ExposeHeaders: []string{
			"Content-Length", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers", "Content-Type",
			"*",
		},

		// Permite credenciales
		AllowCredentials: true,

		// Tiempo máximo de cache para preflight requests
		MaxAge: 24 * time.Hour,
	}

	return cors.New(config)
}
