package user

import (
	"encoding/json"
	"gin-gonic-practice/controller/security"
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/database/redis"
	"gin-gonic-practice/domain/user"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	requestDTO := new(user.User)
	err := c.Bind(requestDTO)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Binding error!",
		})
		return
	}

	if requestDTO.Email == "" || requestDTO.Password == "" {
		c.JSON(400, gin.H{
			"message": "Please input all the values!",
		})
		return
	}

	responseDTO := new(user.User)
	// 레디스에 유저 데이터가 없는 경우
	if redis.IsUserEmpty(requestDTO.Email) == true {
		mariaErr := mariadb.DB.Where("email = ?", requestDTO.Email).Find(&responseDTO).Error
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
	// 레디스 데이터를 가지고 로그인 진행
	userInfo := redis.CLIENT.HGet("users", requestDTO.Email)
	if err := json.Unmarshal([]byte(userInfo.Val()), &responseDTO); err != nil {
		panic(err)
	}

	if responseDTO.Email != requestDTO.Email || responseDTO.Password != requestDTO.Password {
		c.JSON(400, gin.H{
			"message": "Input value error!",
		})
		return
	}

	ts, err := security.CreateToken(requestDTO.Email)
	if err != nil {
		c.JSON(422, gin.H{
			"message": "Status Unprocessable Entity!",
		})
		return
	}

	saveErr := security.CreateAuth(requestDTO.Email, ts)
	if saveErr != nil {
		c.JSON(422, gin.H{
			"message": "Status Unprocessable Entity!",
		})
		return
	}

	c.JSON(200, gin.H{
		"access_token": ts.AccessToken,
		"uuid":         ts.AccessUuid,
		"message":      "Login success!",
	})
	return
}
