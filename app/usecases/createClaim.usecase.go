package usecases

import (
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func CreateClaim(userID, groupID uint, jti string, roles []string) map[string]jwt.MapClaims {

	return map[string]jwt.MapClaims{
		"access_token": {
			"aud":    "http://api.example.com",
			"iss":    "https://krakend.io",
			"sub":    strconv.Itoa(int(userID)),
			"userId": userID,
			"group":  groupID,
			"jti":    jti,
			"roles":  roles,
			"exp":    time.Now().Add(time.Minute).Unix(),
			"kid":    "sim2",
		},
		"refresh_token": {
			"aud":    "http://api.example.com",
			"iss":    "https://krakend.io",
			"sub":    strconv.Itoa(int(userID)),
			"userId": userID,
			"group":  groupID,
			"jti":    jti,
			"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
			"kid":    "sim2",
		},
	}
}
