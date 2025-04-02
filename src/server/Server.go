package server

import (
	database "api-order/src/Database"
	alert "api-order/src/alert/infrastructure/http/routes"
	"api-order/src/client/infrastructure/http/routes"
	"api-order/src/config"
	airData "api-order/src/data/airquality/infrastructure/http/routes"
	lightData "api-order/src/data/light/infrastructure/http/routes"
	motionData "api-order/src/data/motion/infrastructure/http/routes"
	temperatureData "api-order/src/data/temperature/infrastructure/http/routes"
	kit "api-order/src/kit/infrastructure/http/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine   *gin.Engine
	http     string
	port     string
	httpAddr string
}

func NewServer(http, port string) Server {
	gin.SetMode(gin.ReleaseMode)

	srv := Server{
		engine:   gin.New(),
		http:     http,
		port:     port,
		httpAddr: http + ":" + port,
	}

	srv.engine.Use(config.ConfigurationCors())
	database.Connect()
	srv.engine.RedirectTrailingSlash = true
	srv.registerRoutes()

	return srv
}

func (s *Server) registerRoutes() {
	clientRoutes := s.engine.Group("/v1/clients")
	kitRoutes := s.engine.Group("/v1/kits")
	alertRoutes := s.engine.Group("/v1/alerts")
	temperatureRoutes := s.engine.Group("/v1/data/temperature")
	airRoutes := s.engine.Group("/v1/data/airquality")
	lightRoutes := s.engine.Group("/v1/data/light")
	motionRoutes := s.engine.Group("/v1/data/motion")

	kit.KitRoutes(kitRoutes)
	alert.AlertRoutes(alertRoutes)
	routes.ClientRoutes(clientRoutes)
	temperatureData.TemperatureRoutes(temperatureRoutes)
	airData.AirQualityRoutes(airRoutes)
	lightData.LightRoutes(lightRoutes)
	motionData.MotionRoutes(motionRoutes)
}

func (s *Server) Run() {
	log.Println("Server running on " + s.httpAddr)
	s.engine.Run(s.httpAddr)
}
