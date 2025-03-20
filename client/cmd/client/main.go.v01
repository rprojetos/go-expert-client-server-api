package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Sistema iniciado com sucesso!")
	fmt.Println("Iniciando servidor HTTP...")

	// Configuração do cliente HTTP com timeout de 1 segundo
	c := http.Client{Timeout: 1 * time.Second}

	// Corpo da requisição JSON
	jsonData := []byte(`{"moeda":"EUR-BRL"}`)
	url := "http://localhost:8080/cotacao"

	// Executando a requisição POST
	resp, err := c.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	// Lendo e imprimindo a resposta
	log.Println("Status da resposta:", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler a resposta: %v", err)
	}

	log.Println("Resposta do servidor:", string(body))
}
