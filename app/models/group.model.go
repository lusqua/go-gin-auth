package models

import "time"

type Group struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`

	Users []User `gorm:"foreignKey:GroupID"`
}

func (Group) TableName() string {
	return "groups"
}
