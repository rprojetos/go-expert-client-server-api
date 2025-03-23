package storage

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rprojetos/go-expertt/client/internal/config"
	"github.com/rprojetos/go-expertt/client/pkg/fileutil"
)

func SaveQuote(cfg config.Config, result []byte) error {
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
