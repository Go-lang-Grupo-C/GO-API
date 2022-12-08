package main

import (
	"encoding/json"
	"net/http"

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

func handleArticles(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	js, err := json.Marshal(Articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
