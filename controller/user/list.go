package user

import (
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/domain/user"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var userList []user.User
	err := mariadb.DB.Find(&userList).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "유저 목록을 반환할 수 없습니다.",
		})
		return
	}

	c.JSON(200, userList)
}
