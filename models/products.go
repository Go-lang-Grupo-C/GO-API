package models

import (
	"gopkg.in/validator.v2" //importa biblioteca de validação
	"gorm.io/gorm"          //importa biblioteca de banco de dados
)

type Product struct {
	gorm.Model        // o gorm ja traz o id unico e o created_at e o updated_at
	Name       string `json:"name" validate:"nonzero"`
	Price      uint   `json:"price" validate:"nonzero"`
	Code       string `json:"code" validate:"nonzero"`
}

var Products []Product // cria um array de produtos

// a função Validate vai validar os campos do produto e retornar um erro caso nao sejam preenchidos
func Validate(product *Product) error {
	if err := validator.Validate(product); err != nil {
		return err
	}

	return nil
}

