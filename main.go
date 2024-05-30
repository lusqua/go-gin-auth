package main

import (
	"github.com/MicahParks/jwkset"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lusqua/gin-auth/app/config/environment"
	jwtConfig "github.com/lusqua/gin-auth/app/config/jwt"
	"github.com/lusqua/gin-auth/app/controllers/auth"
	"github.com/lusqua/gin-auth/app/models"
	"log"

	"github.com/lusqua/gin-auth/app/config/database"
	"github.com/lusqua/gin-auth/app/controllers/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	database.DbConfig.Connect()
	models.MigrateModels()
	jwtConfig.JWKSetup()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	users.SetUserController(r)
	auth.SetLoginController(r)

	r.GET(
		"/jwk", func(c *gin.Context) {
			rawJWKS, err := jwtConfig.ServerStore.JSONPublic(c)
			if err != nil {
				log.Fatalf("Failed to get the server's JWKS.\nError: %s", err)
			}

			c.JSON(http.StatusOK, rawJWKS)
		},
	)

	r.POST(
		"/login", func(c *gin.Context) {
			token := jwt.New(jwt.SigningMethodRS256)
			token.Header[jwkset.HeaderKID] = jwtConfig.KeyID

			claims := token.Claims.(jwt.MapClaims)
			claims["aud"] = "http://api.example.com"
			claims["iss"] = "https://krakend.io"
			claims["sub"] = "1234567890qwertyuio"
			claims["jti"] = "mnb23vcsrt756yuiomnbvcx98ertyuiop"
			claims["roles"] = []string{"users", "admin"}
			claims["exp"] = 1735689600

			signed, err := token.SignedString(jwtConfig.PrivateKey)
			if err != nil {
				log.Fatalf("Failed to sign a JWT.\nError: %s", err)
			}

			c.JSON(http.StatusOK, gin.H{"token": signed})
		},
	)

	r.Run(":9000")
}
