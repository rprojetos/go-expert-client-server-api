package main

import (
	"log"
	"time"

	"github.com/rprojetos/go-expertt/client/internal/config"
	c "github.com/rprojetos/go-expertt/client/internal/client"
	"github.com/rprojetos/go-expertt/client/internal/storage"
)

func logExecutionTime(startTime time.Time) {
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	log.Println("Tempo total de execução/resposta:", elapsed)
}

func RunQuotesClient (cfg config.Config) {
	startTime := time.Now()

    result, err := c.FetchQuote(cfg)
	if err != nil {
        log.Fatalf("Erro: %v", err)
    }

	storage.SaveQuote(cfg, result)
	if err != nil {
        log.Fatalf("Erro: %v", err)
    }

    logExecutionTime(startTime)

}

func main() {
	log.Println("Carregando configuração do sistema...")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panicf("Erro no carregamento das configurações: %v", err)
	}
	log.Println("Sistema iniciado com sucesso!")

	log.Println("Iniciando client HTTP...")

	RunQuotesClient (cfg)

}
