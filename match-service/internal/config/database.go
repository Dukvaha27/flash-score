package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: os.Getenv("DB_URI"), PreferSimpleProtocol: true}))

	if err != nil {
		panic(err.Error())
	}

	return db
}
