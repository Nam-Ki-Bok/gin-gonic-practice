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

func TestLogout(t *testing.T) {
	var output map[string]string
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
		"Login success!": []byte(`{
    "email": "test@test.com",
    "password": "test"
}`),
	}
	mariadb.Connect()
	redis.Connect()

	router := gin.Default()
	router.POST("/api/v1/user/signup", SignUp)
	router.POST("/api/v1/user/login", Login)
	router.POST("/api/v1/user/logout", Logout)

	// Signup
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/user/signup", bytes.NewBuffer(TestData["Signup success!"]))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}
	assert.Equal(t, "Signup success!", output["message"])

	// Login
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/user/login", bytes.NewBuffer(TestData["Login success!"]))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}
	assert.Equal(t, "Login success!", output["message"])
	testAccessToken := "Bearer " + output["access_token"]

	// Logout
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/user/logout", nil)
	req.Header.Set("Authorization", testAccessToken)
	router.ServeHTTP(w, req)
	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}
	assert.Equal(t, "Logout success!", output["message"])
}
