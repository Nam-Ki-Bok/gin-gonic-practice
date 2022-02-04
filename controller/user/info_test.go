package user

import (
	"bytes"
	"encoding/json"
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/database/redis"
	"gin-gonic-practice/domain/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserInfo(t *testing.T) {
	var output map[string]interface{}
	defer func() {
		mariadb.DB.Where("email = ?", "test@test.com").Delete(&user.User{})
		redis.CLIENT.FlushAll()
	}()

	var TestData = map[string][]byte{
		"Signup success!": []byte(`{
    "email": "test@test.com",
    "password": "test",
    "name": "test"
}`),
		"The user that doesn't exist!": []byte(`{
    "email": "empty@naver.com",
    "password": "test",
	"name": "empty"
}`),
	}
	mariadb.Connect()
	redis.Connect()

	router := gin.Default()
	router.POST("/api/v1/user/signup", SignUp)
	router.GET("/api/v1/user/:email", Info)

	// Signup
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/user/signup", bytes.NewBuffer(TestData["Signup success!"]))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}
	assert.Equal(t, "Signup success!", output["message"])

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "", bytes.NewBuffer(TestData["Signup success!"]))
	req.URL.Path = "/api/v1/user/" + "test@test.com"
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}
	assert.Equal(t, "test", output["name"])
	assert.Equal(t, "Get userInfo success!", output["message"])

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "", bytes.NewBuffer(TestData["The user that doesn't exist!"]))
	req.URL.Path = "/api/v1/user/" + "empty@naver.com"
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}
	assert.Equal(t, "The user that doesn't exist!", output["message"])
}
