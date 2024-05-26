package environment

import (
	_ "github.com/joho/godotenv/autoload"
	"log"

	"fmt"
	"os"
	"reflect"
)

type Env struct {
	HOST     string
	PORT     string
	PASSWORD string
}

var EnvInstance *Env

func init() {
	log.Println("Validating env vars")

	EnvInstance = &Env{
		HOST:     os.Getenv("DB_HOST"),
		PORT:     os.Getenv("DB_PORT"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
	}

	values := reflect.ValueOf(EnvInstance).Elem()

	for i := 0; i < values.NumField(); i++ {
		Name := values.Type().Field(i).Name

		if values.Field(i).String() == "" {
			panic(fmt.Sprintf("Env var %s is not set", Name))
		}
	}

	log.Println("Env vars validated")
}
