package config

import (
	"os"
	"strconv"
)

const (
	DEVELOPER    = "developer"
	HOMOLOGATION = "homologation"
	PRODUCTION   = "production"
)

type Config struct {
	SRV_PORT    string `json:"srv_port"`
	WEB_UI      bool   `json:"web_ui"`
	Mode        string `json:"mode"`
	OpenBrowser bool   `json:"open_browser"`
}

func NewConfig(confi *Config) *Config {
	var conf *Config

	if confi == nil || confi.SRV_PORT == "" {
		conf = defaultConf()
	} else {
		conf = confi
	}

	SRV_PORT := os.Getenv("SRV_PORT")
	if SRV_PORT != "" {
		conf.SRV_PORT = SRV_PORT
	}

	SRV_MODE := os.Getenv("SRV_MODE")
	if SRV_MODE != "" {
		conf.Mode = SRV_MODE
	}

	SRV_WEB_UI := os.Getenv("SRV_WEB_UI")
	if SRV_WEB_UI != "" {
		conf.WEB_UI, _ = strconv.ParseBool(SRV_WEB_UI)
	}

	return conf
}

func defaultConf() *Config {
	Default_conf := Config{
		SRV_PORT:    "8080",
		WEB_UI:      true,
		OpenBrowser: true,
		Mode:        PRODUCTION,
	}

	return &Default_conf
}

// porta que o servidor Gin ira rodar ->
const (
	PortServer = "8080"
)

// configuracoes do banco de dados, strinh de conexão contendo o host, user, password, dbname, port e sslmode
const (
	StringConectDatabase = "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
)

// configuracoes das rotas
const (
	//nome main das rotas/nome da api
	ApiName = "api/v1"
	//nome das rotas
	Products = "products/"
	Product  = "product"

	//nome da rota do usuario
	Users = "user"
)
