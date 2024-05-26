package users

import (
	"github.com/gin-gonic/gin"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
)

func (u *userService) FindUser(userId, groupId uint, userRepo repository.UserRepository) (gin.H, int, error) {

	findUser, err := userRepo.FindUserByIdAndGroup(userId, groupId)

	if err != nil {
		return gin.H{
			"message": "User not found",
		}, 404, err
	}

	return gin.H{
		"user": findUser,
	}, 200, err

}
