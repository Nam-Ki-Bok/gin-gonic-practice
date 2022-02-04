package post

import (
	"bytes"
	"encoding/json"
	"gin-gonic-practice/controller/user"
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/database/redis"
	"gin-gonic-practice/domain/post"
	user2 "gin-gonic-practice/domain/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdatePost(t *testing.T) {
	var output map[string]string
	defer func() {
		mariadb.DB.Where("email = ?", "test@test.com").Delete(&post.Post{})
		mariadb.DB.Where("email = ?", "test@test.com").Delete(&user2.User{})
		redis.CLIENT.FlushAll()
	}()

	mariadb.Connect()
	redis.Connect()

	router := gin.Default()
	router.POST("/api/v1/signup", user.SignUp)
	router.POST("/api/v1/login", user.Login)
	router.POST("/api/v1/post", Create)
	router.PUT("/api/v1/post/:idx", Update)

	// Signup
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/signup", bytes.NewBuffer([]byte(`{
	"email": "test@test.com",
	"password": "test",
	"name": "test"
}`)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}
	assert.Equal(t, "Signup success!", output["message"])

	// Login
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer([]byte(`{
    "email": "test@test.com",
    "password": "test"
}`)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}
	assert.Equal(t, "Login success!", output["message"])
	testAccessToken := "Bearer " + output["access_token"]

	// posting
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/post", bytes.NewBuffer([]byte(`{
	"email": "test@test.com",
	"title": "Test Title",
	"content": "Test Content",
	"name": "test"
}`)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", testAccessToken)
	router.ServeHTTP(w, req)
	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}
	assert.Equal(t, "Posting success!", output["message"])
	idx := output["idx"]

	// Update
	var data = new(post.Post)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "", bytes.NewBuffer([]byte(`{
	"title": "Update Title",
	"content": "Update Content"
}`)))
	req.URL.Path = "/api/v1/post/" + idx
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", testAccessToken)
	router.ServeHTTP(w, req)
	mariadb.DB.Where("idx = ?", idx).Find(&post.Post{}).Scan(data)
	assert.Equal(t, "Update Title", data.Title)
	assert.Equal(t, "Update Content", data.Content)
}
