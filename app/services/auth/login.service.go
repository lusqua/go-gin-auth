package auth

import (
	"github.com/gin-gonic/gin"
	repository "github.com/lusqua/gin-auth/app/repositories/users"
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

	return gin.H{
		"access_token": gin.H{
			"aud":   "http://api.example.com",
			"iss":   "https://krakend.io",
			"sub":   findUser.ID,
			"group": findUser.GroupID,
			"jti":   "mnb23vcsrt756yuiomnbvcx98ertyuiop",
			"roles": []string{"users", "admin"},
			"exp":   1735689600,
		},
		"refresh_token": gin.H{
			"aud":   "http://api.example.com",
			"iss":   "https://krakend.io",
			"sub":   findUser.ID,
			"group": findUser.GroupID,
			"jti":   "mnb23vcsrt756yuiomn12876bvcx98ertyuiop",
			"exp":   1735689600,
		},
		"exp": 1735689600,
	}, nil

}
