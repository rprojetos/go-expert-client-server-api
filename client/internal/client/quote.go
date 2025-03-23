package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/rprojetos/go-expertt/client/internal/config"
)

func FetchQuote(cfg config.Config) ([]byte, error) {
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
