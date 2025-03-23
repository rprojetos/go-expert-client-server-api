package app

import (
	"fmt"
	"log"
	"time"

	"github.com/rprojetos/go-expertt/client/internal/client"
	"github.com/rprojetos/go-expertt/client/internal/config"
	"github.com/rprojetos/go-expertt/client/internal/storage"
)

func RunQuotesClient(cfg config.Config) error {
	startTime := time.Now()

	// Busca cotação
	result, err := client.FetchQuote(cfg)
	if err != nil {
		return fmt.Errorf("falha ao buscar cotação: %w", err)
	}

	// Salva resultado
	if err := storage.SaveQuote(cfg, result); err != nil {
		return fmt.Errorf("falha ao salvar cotação: %w", err)
	}

	// Loga tempo de execução
	logExecutionTime(startTime)
	return nil
}

func logExecutionTime(startTime time.Time) {
	elapsed := time.Since(startTime)
	log.Println("Tempo total de execução:", elapsed)
}
