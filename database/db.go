package database

import (
	"log"

	"github.com/Go-lang-Grupo-C/GO-API/config"
	"github.com/Go-lang-Grupo-C/GO-API/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// função para conectar com o banco de dados
func ConnectDB() {

	dns := config.StringConectDatabase      // var que recebe a string de conexão com o banco de dados
	DB, err = gorm.Open(postgres.Open(dns)) // abre a conexão com o banco de dados

	if err != nil {
		log.Panic("Error to connect database") //traz o erro caso não consiga conectar com o banco de dados
	}

	DB.AutoMigrate(&models.Product{}) //cria a tabela no banco de dados
}
