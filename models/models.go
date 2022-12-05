package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model        // o gorm ja traz o id unico e o created_at e o updated_at
	Name       string `json:"name" validate:"nonzero"`
	Price      uint   `json:"price" validate:"nonzero"`
	Code       string `json:"code" validate:"nonzero"`
}

var Products []Product

func Validade(product *Product) error {
	if err := validator.Validate(product); err != nil {
		return err
	}
	return nil
}
