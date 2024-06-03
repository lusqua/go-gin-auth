package auth

import (
	"github.com/gin-gonic/gin"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
)

var ActiveSessions = make(map[uint]string)

type AuthService interface {
	Login(email, password string, userRepo repository.UserRepository) (gin.H, error)
	Refresh(token string, userRepo repository.UserRepository) (gin.H, error)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}
