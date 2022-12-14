package config

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

// config open browser

const (
	//url do browser
	UrlBrowser = "http://localhost:8080/webui/"
)
