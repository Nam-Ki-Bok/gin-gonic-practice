package main

import (
	"gin-gonic-practice/database/mariadb"
	"gin-gonic-practice/database/redis"
	"gin-gonic-practice/domain/post"
	"gin-gonic-practice/domain/user"
	"gin-gonic-practice/server"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func init() {
	redis.Connect()
	mariadb.Connect()
	router := server.SetupRouter()

	// 레디스에 데이터가 없는 경우 저장
	user.InitRedis(router)
	post.InitRedis(router)

	_ = router.Run(":8080")
}

func main() {
	defer func() {
		err := mariadb.DB.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
}
