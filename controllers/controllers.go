package controllers

import (
	"github.com/Go-lang-Grupo-C/GO-API/models"
	"github.com/gin-gonic/gin"
)

func Products(c *gin.Context) {
	var products []models.Product
	//models.DB.Find(&products)
	c.JSON(200, gin.H{"products": products})

}

func SearchForProduct(c *gin.Context) {
	var product models.Product
	//id := c.Params.ByName("id")
	//database.DB.First(&product, id)
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(200, gin.H{

		"code:":  product.Code,
		"name:":  product.Name,
		"price:": product.Price,
	})
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := models.Validate(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//database.DB.Create(&product)
	c.JSON(200, gin.H{
		"ID":    product.ID,
		"code":  product.Code,
		"name":  product.Name,
		"price": product.Price,
	})

}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	//id := c.Params.ByName("id")
	//database.DB.First(&product, id)
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
	//database.DB.Save(&product)
	c.JSON(200, gin.H{"message": "Product updated is successfully"})
}

func DeleteProduct(c *gin.Context) {
	var product models.Product
	//id := c.Params.ByName("id")
	//database.DB.First(&product, id)
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	//database.DB.Delete(&product)
	c.JSON(200, gin.H{"message": "Product deleted is successfully"})
}
