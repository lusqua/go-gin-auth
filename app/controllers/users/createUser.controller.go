package users

import (
	"github.com/gin-gonic/gin"
	"github.com/lusqua/gin-auth/app/config/database"
	dto "github.com/lusqua/gin-auth/app/dtos/users"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
	service "github.com/lusqua/gin-auth/app/services/users"
)

func CreateUser(c *gin.Context) {

	userDto := dto.NewUserDto()

	database.Connection.AutoMigrate()

	body, err := userDto.CreateUser(c)
	if err != nil {
		return
	}

	userService := service.NewUserService()
	userRepository := repository.NewUserRepository(database.Connection)

	response, _ := userService.CreateUser(body, userRepository)

	if err != nil {
		c.JSON(500, response)
	}

	c.JSON(200, response)

	return
}
