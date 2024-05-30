package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jwtConfig "github.com/lusqua/gin-auth/app/config/jwt"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
	"github.com/lusqua/gin-auth/app/usecases"
)

func (a *authService) Refresh(jti string, userRepo repository.UserRepository) (gin.H, error) {

	userId := ActiveSessions[jti]

	fmt.Println("USER ID: ", userId)

	if userId == 0 {
		fmt.Println("SESSION NOT FOUND")
		return gin.H{
			"message": "invalid token",
		}, nil
	}

	findUser, err := userRepo.FindUserById(userId)

	if err != nil {
		fmt.Println("USER NOT FOUND")
		return gin.H{
			"message": "invalid token",
		}, err
	}

	// delete old session
	delete(ActiveSessions, jti)

	newJti := usecases.GenerateRandomString(32)
	ActiveSessions[newJti] = findUser.ID

	claims := usecases.CreateClaim(findUser.ID, findUser.GroupID, newJti, []string{"user"})

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
