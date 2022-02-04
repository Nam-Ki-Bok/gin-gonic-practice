package user

import (
	"gin-gonic-practice/controller/security"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	security.CheckValidation(c)

	c.JSON(200, gin.H{
		"message": "Logout success!",
	})
	return
}
