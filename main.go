package main

import "github.com/Go-lang-Grupo-C/GO-API/server"

func main() {

	//concet database

	// new sever
	sever := server.NewServer()

	// start server
	sever.Start()
}
