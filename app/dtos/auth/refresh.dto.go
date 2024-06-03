package auth

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (a *authDto) Refresh(c *gin.Context) (string, error) {
	bearerToken := c.GetHeader("Authorization")

	if bearerToken == "" {
		c.JSON(
			401, gin.H{
				"message": "authorization header not found",
			},
		)
		return "", nil
	}

	token := strings.Replace(bearerToken, "Bearer ", "", -1)

	return token, nil
}
