package post

import (
	"gin-gonic-practice/controller/security"
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/database/redis"
	"gin-gonic-practice/domain/post"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	security.CheckValidation(c)

	idx := c.Param("idx")
	redisErr := redis.CLIENT.HGet("posts", idx).Err()
	// 레디스에 게시글 데이터가 없는 경우
	if redisErr == redis.NilError {
		mariaErr := mariadb.DB.Where("idx = ?", idx).Find(&post.Post{}).Error
		// 마리아에 게시글 데이터가 없는 경우
		if mariaErr != nil {
			c.JSON(400, gin.H{
				"message": "The post that doesn't exist!",
			})
			return
		}
	}
	// 마리아 삭제 진행
	mariadb.DB.Delete(&post.Post{}, idx)
	// 레디스 삭제 진행
	redis.CLIENT.HDel("posts", idx)
	c.JSON(200, gin.H{
		"message": "Delete success!",
	})
	return
}
