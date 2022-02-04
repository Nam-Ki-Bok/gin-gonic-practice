package user

import (
	"bytes"
	"encoding/json"
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/database/redis"
	"gin-gonic-practice/domain/user"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUp(t *testing.T) {
	defer func() {
		mariadb.DB.Where("email = ?", "test@test.com").Delete(&user.User{})
		redis.CLIENT.FlushAll()
	}()

	var output map[string]string
	var TestData = map[string][]byte{
		"Please input all the values!": []byte(`{
    "email": "",
    "password": "test",
    "name": "test"
}`),
		"Already using email!": []byte(`{
    "email": "nkb7714@naver.com",
    "password": "test",
    "name": "test"
}`),
		"Signup success!": []byte(`{
    "email": "test@test.com",
    "password": "test",
    "name": "test"
}`),
	}

	mariadb.Connect()
	redis.Connect()
	router := gin.Default()
	router.POST("/api/v1/user/signup", SignUp)

	for message, inputData := range TestData {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/user/signup", bytes.NewBuffer(inputData))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
			panic(err)
		}
		assert.Equal(t, message, output["message"])
	}
}
