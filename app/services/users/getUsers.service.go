package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
	"gorm.io/gorm"
)

func (u *userService) GetUsers(groupId uint, userRepo repository.UserRepository) (gin.H, int, error) {

	findUser, err := userRepo.GetUsersByGroup(groupId)

	if err == nil {
		return gin.H{
			"users": findUser,
		}, 200, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return gin.H{
			"message": "user not found",
		}, 404, err
	}

	return gin.H{}, 500, err

}
