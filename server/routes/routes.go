package routes

import (
	"github.com/Go-lang-Grupo-C/GO-API/config"
	"github.com/Go-lang-Grupo-C/GO-API/controllers"
	"github.com/Go-lang-Grupo-C/GO-API/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) *gin.Engine {

	router.Use(cors.Default())

	main := router.Group(config.ApiName)
	{
		products := main.Group(config.Products)
		{
			products.GET("/", controllers.Products)
			products.GET("/:id", controllers.SearchForProduct)
			products.PATCH("/:id", controllers.UpdateProduct)
			products.POST("/", controllers.CreateProduct)
			products.DELETE("/:id", controllers.DeleteProduct)
		}
		users := main.Group(config.Users)
		{
			users.POST("/login", handler.Login)
		}
	}

	return router
}
