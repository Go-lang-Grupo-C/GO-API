package routes

import (
	"github.com/Go-lang-Grupo-C/GO-API/config"
	"github.com/Go-lang-Grupo-C/GO-API/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group(config.ApiName)
	{
		products := main.Group(config.Products)
		{
			products.POST("/", controllers.CreateProduct)
		}
	}

	return router
}
