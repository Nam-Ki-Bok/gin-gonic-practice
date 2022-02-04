package redis

import (
	"github.com/go-redis/redis/v7"
)

var CLIENT *redis.Client
var NilError interface{}

func Connect() {
	//Initializing redis
	dsn := "localhost:6379"
	client := redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	redisNilError := redis.Nil

	CLIENT = client
	NilError = redisNilError
}
