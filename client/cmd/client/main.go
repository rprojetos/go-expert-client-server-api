package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)


func writeFileString(pathFileName string, content string) int {
	// Abrindo o arquivo com permissão de escrita e criação, se necessário
	// f, err := os.OpenFile(pathFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	f, err := os.OpenFile(pathFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Escrevendo uma string no arquivo e pegando o tamanho
	fileSize, err := f.WriteString(content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Escrita realizada com sucesso no arquivo", f.Name())
	return fileSize
}

func main() {
	fmt.Println("Sistema iniciado com sucesso!")
	fmt.Println("Iniciando servidor HTTP...")

	startTime := time.Now()

	// Cliente HTTP com timeout global de 1 segundo (GuardRail)
	c := http.Client{Timeout: 1 * time.Second}

	jsonData := []byte(`{"moeda":"EUR-BRL"}`)
	url := "http://localhost:8080/cotacao"

	// Cria um contexto com timeout de 300ms
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel() // Libera recursos assim que a função main retornar

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

	log.Println("Resposta do servidor:", string(body))

	fileSize := writeFileString("cotacao.txt", fmt.Sprintf("Dólar: %s", string(body)))
	fmt.Printf("Tamanho do arquivo: %d bytes\n", fileSize)

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Println("Tempo total de execução/resposta:", elapsed)
}
