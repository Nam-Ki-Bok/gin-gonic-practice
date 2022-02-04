package security

import (
	"github.com/gin-gonic/gin"
)

func Unauthorized(c *gin.Context, err error) {
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Unauthorized user!",
		})
		return
	}
}
