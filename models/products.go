package models

import (
	"time"

	"gopkg.in/validator.v2" //importa biblioteca de validação
	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `json:"id" gorm:"primarykey" AutoIncrement:"true"`
	Name      string         `json:"name" validate:"nonzero"`
	Price     uint           `json:"price" validate:"nonzero"`
	Code      string         `json:"code" validate:"nonzero"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

var Products []Product // cria um array de produtos

// a função Validate vai validar os campos do produto e retornar um erro caso nao sejam preenchidos
func Validate(product *Product) error {
	if err := validator.Validate(product); err != nil {
		return err
	}

	return nil
}
