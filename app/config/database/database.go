package database

import (
	"github.com/lusqua/gin-auth/app/config/environment"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var Connection *gorm.DB

var DbConfig = Db{
	Host:     environment.EnvInstance.HOST,
	Port:     environment.EnvInstance.PORT,
	Password: environment.EnvInstance.PASSWORD,
	User:     "postgres",
	DBName:   "postgres",
}

func (c *Db) Connect() {
	dsn := "host=" + c.Host + " user=" + c.User + " password=" + c.Password + " dbname=" + c.DBName + " port=" + c.Port + " sslmode=disable"
	db, err := gorm.Open(
		postgres.Open(dsn), &gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	Connection = db
	println("Connected to database")
}
