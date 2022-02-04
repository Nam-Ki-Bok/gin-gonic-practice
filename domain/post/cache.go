package post

import (
	"encoding/json"
	"gin-gonic-practice/database/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"time"
)

// HSet accepts values in following formats:
//   - HMSet("myhash", "key1", "value1", "key2", "value2")
//   - HMSet("myhash", []string{"key1", "value1", "key2", "value2"})
//   - HMSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})

func InitRedis(router *gin.Engine) {
	isExist, _ := redis.CLIENT.Exists("posts").Uint64()
	if isExist == 0 {
		var postList []Post
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/post/list", nil)
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		if err := json.Unmarshal(w.Body.Bytes(), &postList); err != nil {
			panic(err)
		}

		for _, post := range postList {
			data, err := json.Marshal(post)
			if err != nil {
				panic(err)
			}
			redis.CLIENT.HSet("posts", post.Idx, data)
		}
		redis.CLIENT.Expire("posts", time.Minute*60)
	}
}
