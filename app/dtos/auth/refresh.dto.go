package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lusqua/gin-auth/app/usecases"
)

type RefreshDto struct {
	JTI string `json:"jti"`
}

func (a *authDto) Refresh(c *gin.Context) (RefreshDto, error) {
	var refreshDto RefreshDto

	refreshDto.JTI = c.Param("jti")

	fmt.Println("REFRESH DTO: ", refreshDto.JTI)

	if len(refreshDto.JTI) != 32 {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return RefreshDto{}, fmt.Errorf("Invalid JTI")
	}

	if usecases.ValidateRandomString(refreshDto.JTI) == false {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return RefreshDto{}, fmt.Errorf("Invalid JTI")
	}

	return refreshDto, nil
}
