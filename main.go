package main

import (
	_ "github.com/lusqua/gin-auth/app/config/environment"
	"github.com/lusqua/gin-auth/app/controllers/auth"
	"github.com/lusqua/gin-auth/app/models"

	"fmt"
	"github.com/lusqua/gin-auth/app/config/database"
	"github.com/lusqua/gin-auth/app/controllers/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	database.DbConfig.Connect()
	models.MigrateModels()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	users.SetUserController(r)
	auth.SetLoginController(r)

	//r.POST(
	//	"/auth", func(c *gin.Context) {
	//		c.JSON(
	//			http.StatusOK, gin.H{
	//				"access_token": gin.H{
	//					"aud":   "http://api.example.com",
	//					"iss":   "https://krakend.io",
	//					"sub":   "1234567890qwertyuio",
	//					"jti":   "mnb23vcsrt756yuiomnbvcx98ertyuiop",
	//					"roles": []string{"users", "admin"},
	//					"exp":   1735689600,
	//				},
	//				"refresh_token": gin.H{
	//					"aud": "http://api.example.com",
	//					"iss": "https://krakend.io",
	//					"sub": "1234567890qwertyuio",
	//					"jti": "mnb23vcsrt756yuiomn12876bvcx98ertyuiop",
	//					"exp": 1735689600,
	//				},
	//				"exp": 1735689600,
	//			},
	//		)
	//	},
	//)

	r.GET(
		"/protected/:id", func(c *gin.Context) {
			id := c.Param("id")
			headers := c.Request.Header

			fmt.Println("ID:", id)
			fmt.Println("Headers:", headers)

			c.JSON(http.StatusOK, gin.H{"id": id})
		},
	)

	r.GET(
		"/jwk", func(c *gin.Context) {
			c.JSON(
				http.StatusOK, gin.H{
					"keys": []gin.H{
						{
							"kty": "oct",
							"k":   "AyM1SysPpbyDfgZld3umj1qzKObwVMkoqQ-EstJQLr_T-1qS0gZH75aKtMN3Yj0iPS4hcgUuTwjAzZr1Z9CAow",
							"kid": "sim2",
							"alg": "HS256",
						},
					},
				},
			)
		},
	)

	r.Run(":9000")
}
