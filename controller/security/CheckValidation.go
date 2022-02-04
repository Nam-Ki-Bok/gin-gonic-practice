package security

import (
	"github.com/gin-gonic/gin"
)

func CheckValidation(c *gin.Context) {
	tokenAuth, err := ExtractTokenMetadata(c)
	Unauthorized(c, err)

	_, err = FetchAuth(tokenAuth)
	Unauthorized(c, err)
}
