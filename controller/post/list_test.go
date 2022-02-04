package post

import (
	"encoding/json"
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/domain/post"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostList(t *testing.T) {
	mariadb.Connect()

	var output []map[string]interface{}
	var userList []post.Post
	mariadb.DB.Find(&userList)

	router := gin.Default()
	router.GET("/api/v1/post/list", List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/post/list", nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}
	assert.Equal(t, len(userList), len(output))
}
