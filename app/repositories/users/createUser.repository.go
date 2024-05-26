package users

import (
	"github.com/lusqua/gin-auth/app/models"
	"time"
)

func (ur *userRepository) CreateUser(name, email, password string, isAdmin bool, groupId uint) (models.User, error) {

	newUser := models.User{
		Name:      name,
		Email:     email,
		Password:  password,
		IsAdmin:   isAdmin,
		GroupID:   groupId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	dbOp := ur.db.Create(&newUser)
	if dbOp.Error != nil {
		return models.User{}, dbOp.Error
	}

	return newUser, nil
}
