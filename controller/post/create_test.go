package post

import (
	"bytes"
	"encoding/json"
	userController "gin-gonic-practice/controller/user"
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/database/redis"
	"gin-gonic-practice/domain/post"
	userDB "gin-gonic-practice/domain/user"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	router   = gin.Default()
	output   map[string]string
	TestData = map[string][]byte{
		"Posting success!": []byte(`{
	"email": "test@test.com",
	"title": "Test Title",
	"content": "Test Content",
	"name": "test"
}`),
		"Please input all the values!": []byte(`{
		"email": "test@test.com",
	"title": "",
	"content": "Test Content",
	"name": "test"
}`),
	}
)

func init() {
	mariadb.Connect()
	redis.Connect()

	router.POST("/api/v1/signup", userController.SignUp)
	router.POST("/api/v1/login", userController.Login)
	router.POST("/api/v1/post", Create)
}

func TestCreatePost(t *testing.T) {
	defer func() {
		mariadb.DB.Where("email = ?", "test@test.com").Delete(&post.Post{})
		mariadb.DB.Where("email = ?", "test@test.com").Delete(&userDB.User{})
		redis.CLIENT.FlushAll()
	}()

	signUp(t)
	accessToken := login(t)

	for message, inputData := range TestData {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/post", bytes.NewBuffer(inputData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", accessToken)
		router.ServeHTTP(w, req)

		if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
			panic(err)
		}
		assert.Equal(t, message, output["message"])
	}
}

func signUp(t *testing.T) {
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
}

func login(t *testing.T) string {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer([]byte(`{
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

	return testAccessToken
}
