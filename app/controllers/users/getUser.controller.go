package users

import (
	"github.com/gin-gonic/gin"
	"github.com/lusqua/gin-auth/app/config/database"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
	service "github.com/lusqua/gin-auth/app/services/users"
	"strconv"
)

func GetUser(c *gin.Context) {
	groupId := c.Query("groupId")

	if groupId == "" {
		c.JSON(
			400, gin.H{
				"message": "group id is required",
			},
		)
	}

	uintGroupId, err := strconv.ParseUint(groupId, 10, 32)
	if err != nil {
		return
	}

	userRepo := repository.NewUserRepository(database.Connection)

	userService := service.NewUserService()
	response, statusCode, err := userService.GetUsers(uint(uintGroupId), userRepo)

	c.JSON(statusCode, response)
}
