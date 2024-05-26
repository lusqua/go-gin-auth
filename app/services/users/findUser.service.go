package users

import (
	"github.com/gin-gonic/gin"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
)

func (u *userService) FindUser(userId, groupId uint, userRepo repository.UserRepository) (gin.H, error) {

	findUser, err := userRepo.FindUserByIdAndGroup(userId, groupId)

	if err != nil {
		return gin.H{
			"message": "User not found",
			"error":   err.Error(),
		}, err
	}

	return gin.H{
		"user": findUser,
	}, err

}
