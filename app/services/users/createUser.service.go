package users

import (
	"github.com/gin-gonic/gin"
	dto "github.com/lusqua/gin-auth/app/dtos/users"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
	"golang.org/x/crypto/bcrypt"
)

func (u *userService) CreateUser(body dto.CreateUserDto, userRepo repository.UserRepository) (gin.H, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)

	if err != nil {
		return gin.H{}, err
	}

	user, err := userRepo.CreateUser(body.Name, body.Email, string(bytes), body.IsAdmin, body.GroupID)

	if err != nil {
		return gin.H{}, err
	}

	return gin.H{
		"code": 200,
		"user": user,
	}, nil

}
