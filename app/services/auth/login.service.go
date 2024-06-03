package auth

import (
	"github.com/gin-gonic/gin"
	jwtConfig "github.com/lusqua/gin-auth/app/config/jwt"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
	"github.com/lusqua/gin-auth/app/usecases"
	"golang.org/x/crypto/bcrypt"
)

func (a *authService) Login(email, password string, userRepo repository.UserRepository) (gin.H, error) {

	findUser, err := userRepo.FindUserByEmail(email)
	bytePass := []byte(findUser.Password)

	if err != nil {
		return gin.H{
			"message": "credentials invalid",
		}, err
	}

	err = bcrypt.CompareHashAndPassword(bytePass, []byte(password))

	if err != nil {
		return gin.H{
			"message": "credentials invalid",
		}, err
	}

	jti := usecases.GenerateRandomString(32)
	ActiveSessions[findUser.ID] = jti

	claims := usecases.CreateClaim(findUser.ID, findUser.GroupID, jti, []string{"user"})

	accessTokenString, err := jwtConfig.SignClaim(claims["access_token"])

	if err != nil {
		return gin.H{
			"message": "error signing access token",
		}, err
	}

	refreshTokenString, err := jwtConfig.SignClaim(claims["refresh_token"])

	if err != nil {
		return gin.H{
			"message": "error signing refresh token",
		}, err
	}

	return gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	}, nil
}
