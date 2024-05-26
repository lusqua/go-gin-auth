package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/mail"
)

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *authDto) Login(c *gin.Context) (LoginDto, error) {

	var dto LoginDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": "Unable to parse request body"})
		return LoginDto{}, err
	}

	// Validate Email

	if dto.Email == "" {
		c.JSON(400, gin.H{"error": "Email is required"})
		return LoginDto{}, fmt.Errorf("Email is required")
	}

	if _, err := mail.ParseAddress(dto.Email); err != nil {
		c.JSON(400, gin.H{"error": "Invalid email address"})
		return LoginDto{}, fmt.Errorf("invalid email address")
	}

	// Validate Password

	if dto.Password == "" || len(dto.Password) < 6 {
		c.JSON(400, gin.H{"error": "Password is too short"})
		return LoginDto{}, fmt.Errorf("password is too short")
	}

	return dto, nil
}
