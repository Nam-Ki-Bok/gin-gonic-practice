package user

import (
	"encoding/json"
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/database/redis"
	"gin-gonic-practice/domain/user"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	requestDTO := new(user.User)
	err := c.Bind(requestDTO)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Binding error!",
		})
		return
	}

	if requestDTO.Email == "" || requestDTO.Password == "" || requestDTO.Name == "" {
		c.JSON(400, gin.H{
			"message": "Please input all the values!",
		})
		return
	}

	responseDTO := new(user.User)
	// 레디스에 유저 데이터가 없는 경우
	if redis.IsUserEmpty(requestDTO.Email) == true {
		mariaErr := mariadb.DB.Where("email = ?", requestDTO.Email).Find(&responseDTO).Error
		// 마리아에 유저 데이터가 있는 경우
		if mariaErr == nil {
			// 마리아에 있는 유저 데이터를 레디스에 저장
			input, _ := json.Marshal(responseDTO)
			redis.CLIENT.HSet("users", responseDTO.Email, input)
			c.JSON(400, gin.H{
				"message": "Already using email!",
			})
			return
		}
	} else {
		c.JSON(400, gin.H{
			"message": "Already using email!",
		})
		return
	}

	// 유저 데이터를 마리아에 저장한 뒤 레디스에 저장
	mariadb.DB.Create(requestDTO)
	input, _ := json.Marshal(requestDTO)
	redis.CLIENT.HSet("users", requestDTO.Email, input)
	c.JSON(200, gin.H{
		"email":   requestDTO.Email,
		"message": "Signup success!",
	})
	return
}
