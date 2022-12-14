package handler

import (
	"github.com/Go-lang-Grupo-C/GO-API/models"
	"github.com/gin-gonic/gin"
)

// func login
func Login(c *gin.Context) {
	var user models.User
	var token models.Token

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if user.Username != models.NewAdmin().Username || user.Password != models.NewAdmin().Password {
		c.JSON(401, gin.H{"status": "unauthorized"})
		return
	}

	token.Token = models.USER_TOKEN

	c.Next()
	c.JSON(200, gin.H{"token": token.Token})
}
