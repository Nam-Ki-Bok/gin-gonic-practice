package user

import (
	"encoding/json"
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/database/redis"
	"gin-gonic-practice/domain/user"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	email := c.Param("email")
	responseDTO := new(user.User)

	// 레디스에 유저 데이터가 없는 경우
	if redis.IsUserEmpty(email) == true {
		mariaErr := mariadb.DB.Where("email = ?", email).Find(&responseDTO).Error
		// 마리아에 유저 데이터가 없는 경우
		if mariaErr != nil {
			c.JSON(400, gin.H{
				"message": "The user that doesn't exist!",
			})
			return
		} else {
			// 마리아에 있는 유저 데이터를 레디스에 저장
			input, _ := json.Marshal(responseDTO)
			redis.CLIENT.HSet("users", responseDTO.Email, input)
		}
	}

	// 레디스 데이터 가지고 유저 데이터 반환
	userInfo := redis.CLIENT.HGet("users", email)
	if err := json.Unmarshal([]byte(userInfo.Val()), &responseDTO); err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"idx":     responseDTO.Idx,
		"email":   responseDTO.Email,
		"name":    responseDTO.Name,
		"message": "Get userInfo success!",
	})
	return
}
