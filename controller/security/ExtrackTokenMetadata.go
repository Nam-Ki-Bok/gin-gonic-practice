package security

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ExtractTokenMetadata(c *gin.Context) (*AccessDetails, error) {
	token, err := VerifyToken(c.Request)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		email, ok := claims["email"].(string)
		if !ok {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			Email:      email,
		}, nil
	}
	return nil, err
}
