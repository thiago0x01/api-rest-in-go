package main

import (
	"api-rest-in-go/database"
	"api-rest-in-go/routes"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func loadEnvs() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnvs()

	database.Connect()

	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
