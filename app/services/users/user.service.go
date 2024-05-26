package users

import (
	"github.com/gin-gonic/gin"
	dto "github.com/lusqua/gin-auth/app/dtos/users"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
)

type UserService interface {
	CreateUser(body dto.CreateUserDto, userRepo repository.UserRepository) (gin.H, error)
	GetUsers()
	FindUser()
	UpdateUser()
	DeleteUser()
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}
