package security

import (
	"gin-gonic-practice/database/redis"
)

func FetchAuth(authDetails *AccessDetails) (string, error) {
	email, err := redis.CLIENT.Get(authDetails.AccessUuid).Result()
	if err != nil {
		return "", err
	}

	return email, nil
}
