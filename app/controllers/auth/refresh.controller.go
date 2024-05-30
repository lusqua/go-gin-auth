package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lusqua/gin-auth/app/config/database"
	dto "github.com/lusqua/gin-auth/app/dtos/auth"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
	service "github.com/lusqua/gin-auth/app/services/auth"
)

func Refresh(c *gin.Context) {

	authDto := dto.NewAuthDto()

	refresh, err := authDto.Refresh(c)

	if err != nil {
		return
	}

	authService := service.NewAuthService()
	userRepo := repository.NewUserRepository(database.Connection)

	response, err := authService.Refresh(refresh.JTI, userRepo)

	if err != nil {
		fmt.Println(err)
		c.JSON(401, response)
		return
	}

	fmt.Println("RESPONSE: ", response)

	c.JSON(200, response)

}
