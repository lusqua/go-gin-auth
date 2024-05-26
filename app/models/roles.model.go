package models

type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:120;unique;not null"`
}

func (Role) TableName() string {
	return "roles"
}
