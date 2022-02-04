package post

import (
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/domain/post"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var postList []post.Post
	err := mariadb.DB.Find(&postList).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "게시글 목록을 반환할 수 없습니다.",
		})
		return
	}

	c.JSON(200, postList)
}
