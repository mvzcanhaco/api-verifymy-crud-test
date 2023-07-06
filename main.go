package main

import (
	"fmt"

	"github.com/mvzcanhaco/api-verifymy-crud-test/models"
	"github.com/mvzcanhaco/api-verifymy-crud-test/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")

	models.Users = []models.User{
		{Id: 1, Name: "Nome 1", Age: 12, Email: "teste@test", Password: "112233", Address: "residencia"},
		{Id: 2, Name: "Nome 2", Age: 15, Email: "teste1@test", Password: "112556", Address: "residencia2"},
	}
	routes.HandleResquest()
}
