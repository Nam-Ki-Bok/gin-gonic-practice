package user

import (
	"encoding/json"
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/domain/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserList(t *testing.T) {
	mariadb.Connect()

	var output []map[string]interface{}
	var userList []user.User
	mariadb.DB.Find(&userList)

	router := gin.Default()
	router.GET("/api/v1/user/list", List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/list", nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}
	assert.Equal(t, len(userList), len(output))
}
