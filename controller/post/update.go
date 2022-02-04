package post

import (
	"encoding/json"
	"gin-gonic-practice/controller/security"
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/database/redis"
	"gin-gonic-practice/domain/post"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	security.CheckValidation(c)

	requestDTO := new(post.Post)
	err := c.Bind(requestDTO)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Binding error!",
		})
		return
	}

	idx := c.Param("idx")
	responseDTO := new(post.Post)
	// 레디스에 게시글 데이터가 없는 경우
	redisErr := redis.CLIENT.HGet("posts", idx).Err()
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
			redis.CLIENT.HSet("users", responseDTO.Email, input)
		}
	}
	// 마리아 업데이트 진행
	mariadb.DB.Model(&post.Post{}).Where("idx = ?", idx).Update(requestDTO)

	// 기존 레디스 데이터를 가져온 뒤, 업데이트 후 반영
	tmpDTO := new(post.Post)
	postInfo := redis.CLIENT.HGet("posts", idx)
	if err = json.Unmarshal([]byte(postInfo.Val()), &tmpDTO); err != nil {
		panic(err)
	}
	tmpDTO.Title = requestDTO.Title
	tmpDTO.Content = requestDTO.Content
	input, _ := json.Marshal(tmpDTO)
	redis.CLIENT.HSet("posts", idx, input)

	c.JSON(200, gin.H{
		"message": "Update success!",
	})
	return

}
