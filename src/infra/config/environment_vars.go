package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DatabaseConnectionString = ""
	ApiPort                  = 0
)

func LoadEnv() {
	var err error

	if err := godotenv.Load(); err != nil {
		log.Fatal("Fail on load environment variables")
	}

	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		ApiPort = 3000
	}

	var (
		DATABASE_USER     = os.Getenv("DATABASE_USER")
		DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")
		DATABASE_NAME     = os.Getenv("DATABASE_NAME")
	)

	DatabaseConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		DATABASE_USER,
		DATABASE_PASSWORD,
		DATABASE_NAME,
	)
}
