package main

import (
	"api-rest-in-go/database"
	"api-rest-in-go/routes"
	"fmt"
)

func main() {
	database.Connect()

	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
