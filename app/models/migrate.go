package models

import "github.com/lusqua/gin-auth/app/config/database"

func MigrateModels() {
	var err error

	migrator := database.Connection.Migrator()

	err = migrator.AutoMigrate(&User{})

	if err != nil {
		panic(err)
	}
}
