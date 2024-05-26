package models

type UserRole struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"not null"` // Chave estrangeira para relacionamento com a tabela User
	RoleID uint `gorm:"not null"` // Chave estrangeira para relacionamento com a tabela Role
}

func (UserRole) TableName() string {
	return "user_roles"
}
