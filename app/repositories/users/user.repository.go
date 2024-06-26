package users

import (
	"github.com/lusqua/gin-auth/app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(name, email, password string, isAdmin bool, groupId uint) (models.User, error)
	FindUser(userId uint) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	FindUserById(id uint) (models.User, error)
	FindUserByIdAndGroup(id, group uint) (models.User, error)
	GetUsersByGroup(groupId uint) ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
