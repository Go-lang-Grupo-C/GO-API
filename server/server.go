package server

import (
	"log"
	"net/http"

	"github.com/Go-lang-Grupo-C/GO-API/config"
	"github.com/Go-lang-Grupo-C/GO-API/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	server *gin.Engine
}

// ServeHTTP implements http.Handler
func (*Server) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

func NewServer() *Server {
	return &Server{
		Port:   config.PortServer,
		server: gin.Default(),
	}
}

func (s Server) Start() {
	router := routes.ConfigureRoutes(s.server)
	router.Run(":" + s.Port)
	log.Println("Server running on port: " + s.Port)

}
