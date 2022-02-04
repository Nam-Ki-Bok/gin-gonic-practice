package security

import (
	"gin-gonic-practice/database/redis"
	"time"
)

func CreateAuth(email string, td *Token) error {
	at := time.Unix(td.AtExpires, 0) // converting Unix to UTC
	now := time.Now()

	errAccess := redis.CLIENT.Set(td.AccessUuid, email, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	return nil
}
