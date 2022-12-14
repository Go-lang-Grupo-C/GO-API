package main

import (
	"github.com/Go-lang-Grupo-C/GO-API/config"
	"github.com/Go-lang-Grupo-C/GO-API/database"
	"github.com/Go-lang-Grupo-C/GO-API/server"
	"github.com/Go-lang-Grupo-C/GO-API/server/service"
)

func main() {

	database.ConnectDB()

	// new sever
	sever := server.NewServer()

	// start server

	service.Openbrowser(config.UrlBrowser)
	sever.Start()

}
