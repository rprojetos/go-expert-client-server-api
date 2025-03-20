package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rprojetos/go-expert/internal/handler"
	"github.com/rprojetos/go-expert/internal/bootstrap"
)

func main() {

	err := bootstrap.ConfigSystem()
	if err != nil {
		panic(err)
	}
	fmt.Println("Sistema iniciado com sucesso!")
	fmt.Println("Iniciando servidor HTTP...")

	http.HandleFunc("/cotacao", handler.HandlerCotacaoDolar)

	log.Fatal(http.ListenAndServe(":8080", nil))

	// TODO post ...
	// curl -X POST -H "Content-Type: application/json" -d '{"moeda":"EUR-BRL"}' http://localhost:8080/cotacao
}
