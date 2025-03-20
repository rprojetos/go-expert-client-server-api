package cotacaoapi

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/rprojetos/go-expert/internal/config"
)

func ObterCotacaoDolar() ([]byte, error) {

	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	timeQuoteApi := cfg.Context.Timeout.TimeQuoteApi

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeQuoteApi)*time.Millisecond)
	defer cancel()
	
	req, err := http.NewRequestWithContext(ctx, "GET", cfg.QuoteApiUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %v", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("ERRO: Timeout ao chamar API de cotação - limite de %dms excedido", timeQuoteApi)
			return nil, fmt.Errorf("timeout ao obter cotação do dólar")
		}
		return nil, fmt.Errorf("erro ao chamar API de cotação: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %v", err)
	}
	return body, nil
}
