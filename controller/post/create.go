package post

import (
	"encoding/json"
	"gin-gonic-practice/controller/security"
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/database/redis"
	"gin-gonic-practice/domain/post"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Create(c *gin.Context) {
	security.CheckValidation(c)

	requestDTO := new(post.Post)
	err := c.Bind(requestDTO)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Binding error!",
		})
		return
	}

	if requestDTO.Title == "" {
		c.JSON(400, gin.H{
			"message": "Please input all the values!",
		})
		return
	}

	mariadb.DB.Create(requestDTO)
	input, _ := json.Marshal(requestDTO)
	redis.CLIENT.HSet("posts", requestDTO.Idx, input)

	c.JSON(200, gin.H{
		"idx":     strconv.Itoa(int(requestDTO.Idx)),
		"message": "Posting success!",
	})
	return
}
