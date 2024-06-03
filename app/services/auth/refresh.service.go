package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jwtConfig "github.com/lusqua/gin-auth/app/config/jwt"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
	"github.com/lusqua/gin-auth/app/usecases"
)

func (a *authService) Refresh(token string, userRepo repository.UserRepository) (gin.H, error) {

	claim, err := jwtConfig.GetClaim(token)

	if err != nil {
		return gin.H{
			"message": "invalid token",
		}, nil
	}

	userId := claim["userId"].(float64)
	jti := claim["jti"].(string)

	uintUserId := uint(userId)

	activeSession := ActiveSessions[uintUserId]

	if activeSession != jti {
		fmt.Println("SESSION NOT FOUND")
		return gin.H{
			"message": "invalid token",
		}, nil

	}

	findUser, err := userRepo.FindUserById(uintUserId)

	if err != nil {
		fmt.Println("USER NOT FOUND")
		return gin.H{
			"message": "invalid token",
		}, err
	}

	// delete old session
	delete(ActiveSessions, uintUserId)

	newJti := usecases.GenerateRandomString(32)
	ActiveSessions[uintUserId] = newJti

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
