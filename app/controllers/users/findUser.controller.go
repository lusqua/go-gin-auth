package users

import (
	"github.com/gin-gonic/gin"
	"github.com/lusqua/gin-auth/app/config/database"
	"github.com/lusqua/gin-auth/app/repositories/users"
	"strconv"
)

func FindUser(c *gin.Context) {

	id := c.Param("userId")

	uintId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return
	}

	userRepo := users.NewUserRepository(database.Connection)
	user, err := userRepo.FindUserById(uint(uintId))

	c.JSON(
		200, gin.H{
			"user": user,
		},
	)

}
