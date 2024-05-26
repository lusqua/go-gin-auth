package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/mail"
)

type CreateUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	IsAdmin  bool   `json:"isAdmin"`
	GroupID  uint   `json:"groupId"`
}

func (u *userDto) CreateUser(c *gin.Context) (CreateUserDto, error) {

	var dto CreateUserDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": "Unable to parse request body"})
		return CreateUserDto{}, err
	}

	// Validate Email

	if dto.Email == "" {
		c.JSON(400, gin.H{"error": "Email is required"})
		return CreateUserDto{}, fmt.Errorf("Email is required")
	}

	if _, err := mail.ParseAddress(dto.Email); err != nil {
		c.JSON(400, gin.H{"error": "Invalid email address"})
		return CreateUserDto{}, fmt.Errorf("invalid email address")
	}

	// Validate Password

	if dto.Password == "" || len(dto.Password) < 6 {
		c.JSON(400, gin.H{"error": "Password is required and must be at least 6 characters"})
		return CreateUserDto{}, fmt.Errorf("password is required and must be at least 6 characters")
	}

	// Validate Name

	if dto.Name == "" {
		c.JSON(400, gin.H{"error": "Name is required"})
		return CreateUserDto{}, fmt.Errorf("name is required")
	}

	// Validate GroupID

	if dto.GroupID == 0 {
		c.JSON(400, gin.H{"error": "GroupID is required"})
		return CreateUserDto{}, fmt.Errorf("GroupID is required")
	}

	return dto, nil
}
