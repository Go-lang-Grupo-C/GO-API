package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model        // o gorm ja traz o id unico e o created_at e o updated_at
	Name       string `json:"name"`
	Price      uint   `json:"price"`
	Code       string `json:"code"`
}

var Products []Product
