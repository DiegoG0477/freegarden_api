// src/server/Server.go
package server

import (
	database "api-order/src/Database"
	alertRoutes "api-order/src/alert/infrastructure/http/routes" // Alias si es necesario
	"api-order/src/config"
	dataRoutes "api-order/src/gardendata/infrastructure/http/routes"
	kitRoutes "api-order/src/kit/infrastructure/http/routes"
	userRoutes "api-order/src/user/infrastructure/http/routes"
	"log"

	"github.com/gin-gonic/gin"

	// ¡IMPORTANTE! Importa los documentos generados por swag
	// Reemplaza 'api-order' con el nombre real de tu módulo en go.mod
	_ "api-order/src/docs" // El _ indica que solo se usa por sus efectos secundarios (init)

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Server struct {
	engine   *gin.Engine
	http     string
	port     string
	httpAddr string
}

func NewServer(http, port string) Server {
	gin.SetMode(gin.ReleaseMode) // O gin.DebugMode durante el desarrollo

	srv := Server{
		engine:   gin.New(), // Considera gin.Default() si quieres los middlewares por defecto (Logger, Recovery)
		http:     http,
		port:     port,
		httpAddr: http + ":" + port,
	}

	// Middlewares
	if gin.Mode() == gin.DebugMode {
		srv.engine.Use(gin.Logger()) // Añadir logger en modo debug
	}
	srv.engine.Use(gin.Recovery()) // Añadir recovery para panics
	srv.engine.Use(config.ConfigurationCors())
	database.Connect()
	srv.engine.RedirectTrailingSlash = true
	srv.registerRoutes()

	return srv
}

func (s *Server) registerRoutes() {
	// Ruta para Swagger UI
	// Asegúrate que el BasePath ('/v1' en este caso) no interfiera.
	// Sirviendo Swagger fuera del grupo /v1 es común.
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Swagger UI available at /swagger/index.html")

	// Grupos de rutas v1
	v1 := s.engine.Group("/v1") // Agrupa todas tus rutas bajo /v1

	userRoutesGroup := v1.Group("/users")
	kitRoutesGroup := v1.Group("/kits")
	alertRoutesGroup := v1.Group("/alerts")
	dataRoutesGroup := v1.Group("/garden/data")

	kitRoutes.KitRoutes(kitRoutesGroup)
	alertRoutes.AlertRoutes(alertRoutesGroup)
	userRoutes.UserRoutes(userRoutesGroup)
	dataRoutes.GardenDataRoutes(dataRoutesGroup)

}

func (s *Server) Run() {
	log.Println("Server running on " + s.httpAddr)
	// Usa ListenAndServe para manejar errores de inicio
	if err := s.engine.Run(s.httpAddr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
