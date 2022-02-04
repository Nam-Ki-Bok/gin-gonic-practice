package security

import (
	"gin-gonic-practice/database/redis"
)

func DeleteAuth(accessUuid string) (int64, error) {
	deleted, err := redis.CLIENT.Del(accessUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
