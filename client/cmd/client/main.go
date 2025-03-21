package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/rprojetos/go-expertt/client/internal/config"
	"github.com/rprojetos/go-expertt/client/pkg/fileutil"
)

func loadConfig() (config.Config, error){
	cfg, err := config.LoadConfig()
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func fetchCotacao(cfg config.Config) ([]byte, error){
	// Cliente HTTP com timeout global de 1 segundo (GuardRail)
	c := http.Client{Timeout: 1 * time.Second}

	timeResponseApi := cfg.Context.Timeout.TimeResponseApi

	// Cria um contexto com timeout de 300ms
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeResponseApi)*time.Millisecond)
	defer cancel() // Libera recursos assim que a função main retornar

	jsonData := []byte(`{"moeda":"EUR-BRL"}`)
	url := cfg.Cotacao.Url

	// Cria a requisição com o contexto
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	// Define o header Content-Type
	req.Header.Set("Content-Type", "application/json")

	// Executa a requisição
	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição: %w", err)
	}
	defer resp.Body.Close()

	log.Println("Status da resposta:", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler a resposta: %w", err)
	}

	return body, nil
}

func saveResult(cfg config.Config, result []byte) error {
	var resultado struct {
		Bid string `json:"bid"`
	}

	// Faz o parse do JSON
	if err := json.Unmarshal(result, &resultado); err != nil {
		return fmt.Errorf("erro ao processar dados: %w", err)
	}

	log.Println("Valor do BID:", resultado.Bid)

	pathFileName := cfg.Cotacao.PathFileName
	fileSize, err := fileutil.WriteFileString(pathFileName, fmt.Sprintf("Dólar: %s\n", resultado.Bid))
	log.Printf("Tamanho do arquivo: %s >> %d bytes\n", pathFileName, fileSize)
	if err != nil {
		return fmt.Errorf("erro ao processar dados: %w", err)
	}
	return nil
}

func logExecutionTime(startTime time.Time) {
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	log.Println("Tempo total de execução/resposta:", elapsed)
}

func RunQuotesClient (cfg config.Config) {
	startTime := time.Now()

    result, err := fetchCotacao(cfg)
	if err != nil {
        log.Fatalf("Erro: %v", err)
    }

	saveResult(cfg, result)
	if err != nil {
        log.Fatalf("Erro: %v", err)
    }

    logExecutionTime(startTime)

}

func main() {
	log.Println("Carregando configuração do sistema...")
	cfg, err := loadConfig()
	if err != nil {
		log.Panicf("Erro no carregamento das configurações: %v", err)
	}
	log.Println("Sistema iniciado com sucesso!")

	log.Println("Iniciando client HTTP...")

	RunQuotesClient (cfg)

}
