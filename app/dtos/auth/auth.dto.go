package auth

import "github.com/gin-gonic/gin"

type AuthDto interface {
	Login(c *gin.Context) (LoginDto, error)
}

type authDto struct {
}

func NewAuthDto() AuthDto {
	return &authDto{}
}
