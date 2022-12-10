package routes

import (
	"github.com/Go-lang-Grupo-C/GO-API/config"
	"github.com/Go-lang-Grupo-C/GO-API/controllers"
	"github.com/Go-lang-Grupo-C/GO-API/handler"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group(config.ApiName)
	{
		products := main.Group(config.Products)
		{
			products.GET("/", controllers.Products)

		}
		product := main.Group(config.Product)
		{
			product.GET("/:id", controllers.SearchForProduct)
			product.PATCH("/:id", controllers.UpdateProduct)
			product.POST("/", controllers.CreateProduct)
			product.DELETE("/:id", controllers.DeleteProduct)

		}
		users := main.Group(config.Users)
		{
			users.POST("/login", handler.Login)

		}
	}

	return router
}
