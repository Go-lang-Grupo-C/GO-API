package server

import (
	"log"

	"github.com/Go-lang-Grupo-C/GO-API/config"
	"github.com/Go-lang-Grupo-C/GO-API/server/routes"
	"github.com/Go-lang-Grupo-C/GO-API/webui"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	server *gin.Engine
}

func NewServer() *Server {
	return &Server{
		Port:   config.PortServer,
		server: gin.Default(),
	}
}

func (s Server) Start() {
	router := routes.ConfigureRoutes(s.server)
	webui.RegisterUIHandlers(router)
	router.Run(":" + s.Port)
	log.Println("Server running on port: " + s.Port)

}
