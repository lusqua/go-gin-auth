package users

import "github.com/gin-gonic/gin"

type UserDto interface {
	CreateUser(c *gin.Context) (CreateUserDto, error)
}

type userDto struct {
}

func NewUserDto() UserDto {
	return &userDto{}
}
