package auth

import "github.com/gin-gonic/gin"

type AuthDto interface {
	Login(c *gin.Context) (LoginDto, error)
	Refresh(c *gin.Context) (string, error)
}

type authDto struct {
}

func NewAuthDto() AuthDto {
	return &authDto{}
}
