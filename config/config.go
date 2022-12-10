package config

//configuracoes do server
const (
	//Port is the port that the server will run on
	PortServer = "8080"
)

//configuracoes do banco de dados
const (
	StringConectDatabase = "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
)

//configuracoes das rotas
const (
	//api name
	ApiName = "api/v1"
	//products
	Products = "products/"
	Product  = "product"

	//users
	Users = "user"
)
