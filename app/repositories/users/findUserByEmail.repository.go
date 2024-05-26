package users

import "github.com/lusqua/gin-auth/app/models"

func (ur *userRepository) FindUserByEmail(email string) (models.User, error) {

	var user models.User

	dbOp := ur.db.Model(&models.User{}).Where("email = ?", email).First(&user)

	if dbOp.Error != nil {
		return models.User{}, dbOp.Error
	}

	return user, nil
}
