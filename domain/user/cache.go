package user

import (
	"encoding/json"
	"gin-gonic-practice/database/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"time"
)

func InitRedis(router *gin.Engine) {
	isExist, _ := redis.CLIENT.Exists("users").Uint64()
	if isExist == 0 {
		var userList []User
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/user/list", nil)
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		if err := json.Unmarshal(w.Body.Bytes(), &userList); err != nil {
			panic(err)
		}

		for _, user := range userList {
			data, err := json.Marshal(user)
			if err != nil {
				panic(err)
			}
			redis.CLIENT.HSet("users", user.Email, data)
		}
		redis.CLIENT.Expire("users", time.Minute*60)
	}
}
