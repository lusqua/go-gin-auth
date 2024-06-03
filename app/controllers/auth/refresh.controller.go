package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lusqua/gin-auth/app/config/database"
	dtos "github.com/lusqua/gin-auth/app/dtos/auth"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
	service "github.com/lusqua/gin-auth/app/services/auth"
)

func Refresh(c *gin.Context) {

	authDto := dtos.NewAuthDto()
	token, err := authDto.Refresh(c)

	if err != nil {
		return
	}

	authService := service.NewAuthService()
	userRepo := repository.NewUserRepository(database.Connection)

	response, err := authService.Refresh(token, userRepo)

	if err != nil {
		fmt.Println(err)
		c.JSON(401, response)
		return
	}

	c.JSON(200, response)
}
