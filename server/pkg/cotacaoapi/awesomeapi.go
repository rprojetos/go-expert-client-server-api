package cotacaoapi

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func ObterCotacaoDolar() ([]byte, error) {

	apiURL := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %v", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("ERRO: Timeout ao chamar API de cotação - limite de 200ms excedido")
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
