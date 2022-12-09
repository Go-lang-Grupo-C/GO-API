package controllers

import (
	"strconv"

	"github.com/Go-lang-Grupo-C/GO-API/database"
	"github.com/Go-lang-Grupo-C/GO-API/models"
	"github.com/gin-gonic/gin"
)

// função que lista todos os produtos
func Products(c *gin.Context) {

	var products []models.Product
	database.DB.Find(&products)
	c.JSON(200, gin.H{"list": products})

}

// função que busca somente um produto por ID via paramentro
func SearchForProduct(c *gin.Context) {

	var product models.Product
	id := c.Param("id")
	newId, _ := strconv.Atoi(id)

	database.DB.First(&product, newId)

	//valido se o prodtuo existe caso ele for id 0 retorna um erro
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(200, gin.H{
		//me tras apenas codigo, nome e preço
		"code:":  product.Code,
		"name:":  product.Name,
		"price:": product.Price,
	})
}

// criando novo produto
func CreateProduct(c *gin.Context) {

	var product models.Product
	//valido se o produto foi preenchido corretamente via corpo da requisição JSON
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//chamo o metodo de validação cirado no models
	if err := models.Validate(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&product)

	//retorno o produto criado
	c.JSON(200, gin.H{
		"ID":    product.ID,
		"code":  product.Code,
		"name":  product.Name,
		"price": product.Price,
	})

}

// editar o produto pego ele via paramentro pelo ID
func UpdateProduct(c *gin.Context) {

	var product models.Product
	id := c.Params.ByName("id")
	database.DB.First(&product, id)

	//valido se o produto existe
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	//valido se o produto foi preenchido corretamente via corpo da requisição JSON
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//chamo o metodo de validação cirado no models
	if err := models.Validate(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&product)

	//retorno o mensagem de sucesso
	c.JSON(200, gin.H{"message": "Product updated is successfully"})
}

// função de deletar produto
func DeleteProduct(c *gin.Context) {

	var product models.Product
	id := c.Params.ByName("id")
	database.DB.First(&product, id)

	//valido se o produto existe
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	database.DB.Delete(&product)

	//retorno o mensagem de sucesso
	c.JSON(200, gin.H{"message": "Product deleted is successfully"})
}
