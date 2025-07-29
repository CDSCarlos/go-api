package main

import (
	"fmt"
	"go-api/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.RegisterRoutes()
	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
