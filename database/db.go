package database

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() {
	var envs map[string]string

	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := envs["CONNECTIONSTRING"]

	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}
}
