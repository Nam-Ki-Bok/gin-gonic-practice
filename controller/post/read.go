package post

import (
	"encoding/json"
	"gin-gonic-practice/controller/security"
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/database/redis"
	"gin-gonic-practice/domain/post"
	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {
	security.CheckValidation(c)

	idx := c.Param("idx")
	responseDTO := new(post.Post)
	redisErr := redis.CLIENT.HGet("posts", idx).Err()
	// 레디스에 게시글 데이터가 없는 경우
	if redisErr == redis.NilError {
		mariaErr := mariadb.DB.Where("idx = ?", idx).Find(&responseDTO).Error
		// 마리아에 게시글 데이터가 없는 경우
		if mariaErr != nil {
			c.JSON(400, gin.H{
				"message": "The post that doesn't exist!",
			})
			return
		} else {
			// 마리아에 있는 게시글 데이터를 레디스에 저장
			input, _ := json.Marshal(responseDTO)
			redis.CLIENT.HSet("posts", responseDTO.Idx, input)
		}
	}
	// 레디스 데이터를 가지고 읽기 진행
	postInfo := redis.CLIENT.HGet("posts", idx)
	if err := json.Unmarshal([]byte(postInfo.Val()), &responseDTO); err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"email":   responseDTO.Email,
		"title":   responseDTO.Title,
		"content": responseDTO.Content,
		"name":    responseDTO.Name,
	})
	return
}
