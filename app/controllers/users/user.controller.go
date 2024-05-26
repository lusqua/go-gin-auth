package users

import "github.com/gin-gonic/gin"

func SetUserController(r *gin.Engine) {
	r.POST("/users", CreateUser)
	r.GET("/users/:userId", FindUser)
}
