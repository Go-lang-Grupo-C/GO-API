package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/Go-lang-Grupo-C/GO-API/config"
	"github.com/Go-lang-Grupo-C/GO-API/database"
	"github.com/Go-lang-Grupo-C/GO-API/server"
)

func main() {

	Default_conf := &config.Config{}

	if file_config := os.Getenv("STOQ_CONFIG"); file_config != "" {
		file, _ := os.ReadFile(file_config)
		_ = json.Unmarshal(file, &Default_conf)
	}

	conf := config.NewConfig(Default_conf)

	database.ConnectDB()

	// new sever
	sever := server.NewServer()

	// start server
	//sever.Start()

	done := make(chan bool)
	go sever.Start()
	log.Printf("Server Run on Port: %v, Mode: %v,WEBUI: %v", conf.SRV_PORT, conf.Mode, conf.WEB_UI)
	open(fmt.Sprintf("http://localhost:%v", conf.SRV_PORT), conf)
	<-done
}

func open(url string, conf *config.Config) error {
	var cmd string
	var args []string

	if !conf.OpenBrowser {
		return nil
	}

	switch runtime.GOOS {
	case "windows": // For Windows
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin": // Mac OS
		cmd = "open"
	default: // For "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
