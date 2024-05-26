package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lusqua/gin-auth/app/config/database"
	dto "github.com/lusqua/gin-auth/app/dtos/auth"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
	svc "github.com/lusqua/gin-auth/app/services/auth"
)

func Login(c *gin.Context) {

	authDto := dto.NewAuthDto()

	body, err := authDto.Login(c)
	if err != nil {
		return
	}

	userRepo := repository.NewUserRepository(database.Connection)

	authService := svc.NewAuthService()
	response, err := authService.Login(body.Email, body.Password, userRepo)

	if err != nil {
		c.JSON(401, response)
		return
	}

	c.JSON(200, response)
}
