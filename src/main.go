// src/main.go
package main

import (
	"log"
	"os"

	"api-order/src/server" // Asegúrate que la ruta del módulo sea correcta

	"github.com/joho/godotenv"
)

// @title           API Hexagonal Go (Sensor Kits)
// @version         1.0
// @description     API para gestionar kits de sensores y sus datos.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.tu-soporte.com/support
// @contact.email  support@tu-dominio.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080  // Cambia esto por tu host y puerto reales (puedes usar variables de entorno)
// @BasePath  /v1             // Base path de tu API v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Reemplaza localhost:8080 en @host si usas variables de entorno
	HOST := os.Getenv("HOST_SERVER")
	PORT := os.Getenv("PORT_SERVER")
	// swaggerHost := HOST + ":" + PORT // Puedes usar esto para el @host

	if HOST == "" || PORT == "" {
		log.Fatal("HOST_SERVER or PORT_SERVER is not set")
	}

	srv := server.NewServer(HOST, PORT)
	srv.Run()
}
