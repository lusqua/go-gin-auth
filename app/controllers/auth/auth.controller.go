package auth

import "github.com/gin-gonic/gin"

func SetLoginController(r *gin.Engine) {

	authGroup := r.Group("/auth")

	authGroup.POST("/login", Login)
	authGroup.POST("/refresh", Refresh)
}
