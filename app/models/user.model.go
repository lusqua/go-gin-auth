package models

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:120;not null"`
	Email    string `gorm:"size:120;unique;not null;index"`
	Password string `gorm:"size:256;not null"`
	IsAdmin  bool   `gorm:"default:false"`

	GroupID   uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`

	Group *Group `gorm:"foreignKey:GroupID;not null"`
	Roles []Role `gorm:"many2many:user_roles;"`
}

func (User) TableName() string {
	return "users"
}
