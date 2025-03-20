package main

import (
	"fmt"
	"log"
	"net/http"

    "github.com/rprojetos/go-expert/internal/handler"
)


func main() {
	fmt.Println("Iniciando servidor HTTP...")

	http.HandleFunc("/cotacao", handler.HandlerCotacaoDolar)

	log.Fatal(http.ListenAndServe(":8080", nil))

	// TODO post ...
	// curl -X POST -H "Content-Type: application/json" -d '{"moeda":"EUR-BRL"}' http://localhost:8080/cotacao
}
