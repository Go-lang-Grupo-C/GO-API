package handler

import (
	"github.com/Go-lang-Grupo-C/GO-API/models"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) bool {

	token := c.GetHeader("Authorization")

	if len(token) > 0 && (token[7:] == models.USER_TOKEN) {
		return true
	} else {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		return false
	}

}
