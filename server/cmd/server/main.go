package main

import (
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
	log.Println("Sistema iniciado com sucesso!")
	log.Println("Iniciando servidor HTTP...")

	http.HandleFunc("/cotacao", handler.HandlerCotacaoDolar)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
