package controllers

import (
	"github.com/Go-lang-Grupo-C/GO-API/database"
	"github.com/Go-lang-Grupo-C/GO-API/handler"
	"github.com/Go-lang-Grupo-C/GO-API/models"
	"github.com/gin-gonic/gin"
)

// função que lista todos os produtos
func Products(c *gin.Context) {

	if !handler.Auth(c) {
		return
	}

	var products []models.Product
	database.DB.Order("id").Find(&products)
	c.JSON(200, gin.H{"list": products})

}

// função que busca somente um produto por ID via paramentro
func SearchForProduct(c *gin.Context) {

	if !handler.Auth(c) {
		return
	}

	var product models.Product
	id := c.Param("id")
	database.DB.First(&product, id)
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(200, product)

}

// criando novo produto
func CreateProduct(c *gin.Context) {

	if !handler.Auth(c) {
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := models.Validate(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&product)
	c.JSON(200, gin.H{
		"ID":    product.ID,
		"code":  product.Code,
		"name":  product.Name,
		"price": product.Price,
	})

}

// editar o produto pego ele via paramentro pelo ID
func UpdateProduct(c *gin.Context) {

	if !handler.Auth(c) {
		return
	}

	var product models.Product
	id := c.Params.ByName("id")
	database.DB.First(&product, id)
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := models.Validate(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&product)
	c.JSON(200, gin.H{"message": "Product updated is successfully"})

}

// função de deletar produto
func DeleteProduct(c *gin.Context) {

	if !handler.Auth(c) {
		return
	}

	var product models.Product
	id := c.Param("id")
	database.DB.First(&product, id)
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	database.DB.Delete(&product)
	c.JSON(200, gin.H{"message": "Product deleted is successfully"})

}
