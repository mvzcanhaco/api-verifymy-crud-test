package main

import (
	"fmt"
	"time"

	"github.com/mvzcanhaco/api-verifymy-crud-test/database"
	"github.com/mvzcanhaco/api-verifymy-crud-test/routes"
)

func main() {
	time.Sleep(5 * time.Second) // Atraso de 5 segundos
	database.ConnectDB()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
