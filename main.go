package main

import (
	"github.com/Go-lang-Grupo-C/GO-API/database"
	"github.com/Go-lang-Grupo-C/GO-API/server"
	"github.com/Go-lang-Grupo-C/GO-API/server/service"
)

func main() {

	database.ConnectDB()

	// new sever
	sever := server.NewServer()

	// start server

	service.Openbrowser("http://localhost:8080/webui/")
	sever.Start()

}
