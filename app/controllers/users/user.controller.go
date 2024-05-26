package users

import "github.com/gin-gonic/gin"

func SetUserController(r *gin.Engine) {
	userGroup := r.Group("/users")

	userGroup.POST("/", CreateUser)
	userGroup.GET("/", GetUser)
	userGroup.GET("/:userId", FindUser)
}
