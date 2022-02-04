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

func TestLogin(t *testing.T) {
	defer func() {
		mariadb.DB.Where("email = ?", "test@test.com").Delete(&user.User{})
		redis.CLIENT.FlushAll()
	}()

	var output map[string]string
	var TestData = map[string][]byte{
		"Login success!": []byte(`{
    "email": "test@test.com",
    "password": "test",
	"name": "test"
}`),
		"Input value error!": []byte(`{
	"email": "test@test.com",
	"password": "wrong password"
}`),
		"Please input all the values!": []byte(`{
	"email": "",
	"password": "test"
}`),
		"The user that doesn't exist!": []byte(`{
	"email": "anonymous",
	"password": "test"
}`),
	}

	mariadb.Connect()
	redis.Connect()

	router := gin.Default()

	router.POST("/api/v1/user/signup", SignUp)
	router.POST("/api/v1/user/login", Login)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/user/signup", bytes.NewBuffer(TestData["Login success!"]))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}
	assert.Equal(t, "Signup success!", output["message"])

	for message, inputData := range TestData {
		w = httptest.NewRecorder()
		req, _ = http.NewRequest("POST", "/api/v1/user/login", bytes.NewBuffer(inputData))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
			panic(err)
		}
		assert.Equal(t, message, output["message"])
	}
}
