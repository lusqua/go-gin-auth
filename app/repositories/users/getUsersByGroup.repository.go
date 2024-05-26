package users

import "github.com/lusqua/gin-auth/app/models"

func (ur *userRepository) GetUsersByGroup(groupId uint) ([]models.User, error) {

	var users []models.User

	dbOp := ur.db.Model(&models.User{}).Where("group_id = ?", groupId).Find(&users)

	if dbOp.Error != nil {
		return users, dbOp.Error
	}

	return users, nil
}
