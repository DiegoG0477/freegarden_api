package server

import (
	database "api-order/src/Database"
	"api-order/src/client/infraestructure/http/routes"
	order "api-order/src/order/infraestructure/http/routes"
	product "api-order/src/product/infraestructure/http/routes"
	"log"
	"api-order/src/config" 
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine 		*gin.Engine
	http 		string
	port 		string
	httpAddr 	string
}

func NewServer(http, port string) Server {
	gin.SetMode(gin.ReleaseMode)

	srv := Server{
		engine: gin.New(),
		http: http,
		port: port,
		httpAddr: http + ":" + port,
	}

	srv.engine.Use(config.ConfigurationCors())
	database.Connect()
	srv.engine.RedirectTrailingSlash = true
	srv.registerRoutes()

	

	return srv
}

func (s *Server) registerRoutes(){
	clientRoutes := s.engine.Group("/v1/client")
	orderRoutes := s.engine.Group("/v1/order")
	productRoutes := s.engine.Group("/v1/product")

	routes.ClientRoutes(clientRoutes)
	order.OrderRoutes(orderRoutes)
	product.ProductRoutes(productRoutes)
}

func (s *Server) Run() {
	log.Println("Server running on " + s.httpAddr)
	s.engine.Run(s.httpAddr)
}
