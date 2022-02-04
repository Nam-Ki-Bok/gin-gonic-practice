package server

import (
	"gin-gonic-practice/controller/post"
	"gin-gonic-practice/controller/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userAPI := router.Group("/api/v1/user")
	{
		userAPI.POST("/signup", user.SignUp)
		userAPI.POST("/login", user.Login)
		userAPI.POST("/logout", user.Logout)

		userAPI.GET("/list", user.List)
		userAPI.GET("/:email", user.Info)
	}

	postAPI := router.Group("/api/v1/post")
	{
		postAPI.POST("", post.Create)
		postAPI.GET("/:idx", post.Read)
		postAPI.PUT("/:idx", post.Update)
		postAPI.DELETE("/:idx", post.Delete)

		postAPI.GET("/list", post.List)
	}

	return router
}
