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


func main() {
	fmt.Println("Sistema iniciado com sucesso!")
	fmt.Println("Iniciando servidor HTTP...")

	cfg, err := config.LoadConfig()
	if err != nil {
		// log.Fatalf("Erro no carregamento das configurações: %v", err)
		log.Panicf("Erro no carregamento das configurações: %v", err)
	}

	startTime := time.Now()

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
		log.Fatalf("Erro ao criar a requisição: %v", err)
	}

	// Define o header Content-Type
	req.Header.Set("Content-Type", "application/json")

	// Executa a requisição
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	log.Println("Status da resposta:", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler a resposta: %v", err)
	}

	var resultado struct {
		Bid string `json:"bid"`
	}

	// Faz o parse do JSON
	if err := json.Unmarshal(body, &resultado); err != nil {
		log.Fatalf("Erro ao processar dados: %v", err)
	}

	log.Println("Resposta do servidor:", resultado.Bid)

	pathFileName := cfg.Cotacao.PathFileName
	fileSize := fileutil.WriteFileString(pathFileName, fmt.Sprintf("Dólar: %s\n", resultado.Bid))
	fmt.Printf("Tamanho do arquivo: %d bytes\n", fileSize)

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Println("Tempo total de execução/resposta:", elapsed)
}
