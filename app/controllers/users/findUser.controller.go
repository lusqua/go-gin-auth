package users

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/lusqua/gin-auth/app/config/database"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
	service "github.com/lusqua/gin-auth/app/services/users"
	"strconv"
)

func FindUser(c *gin.Context) {

	id := c.Param("userId")
	groupId := c.Query("groupId")

	if groupId == "" {
		c.JSON(
			400, gin.H{
				"message": "group id is required",
			},
		)
	}

	uintId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return
	}

	uintGroupId, err := strconv.ParseUint(groupId, 10, 32)
	if err != nil {
		return
	}

	userRepo := repository.NewUserRepository(database.Connection)

	userService := service.NewUserService()
	response, err := userService.FindUser(uint(uintId), uint(uintGroupId), userRepo)

	c.JSON(200, response)

}
