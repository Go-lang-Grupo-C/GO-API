package main

import (
	"github.com/Go-lang-Grupo-C/GO-API/database"
	"github.com/Go-lang-Grupo-C/GO-API/server"
)

func main() {

	database.ConnectDB()

	// new sever
	sever := server.NewServer()

	// start server
	sever.Start()
}
