package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type CotacaoDolar struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func obterCotacaoDolar() ([]byte, error) {

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

func handlerCotacaoDolar(w http.ResponseWriter, r *http.Request) {
	dadosCotacao, err := obterCotacaoDolar()
	if err != nil {
		http.Error(w, fmt.Sprintf("Falha ao obter cotação: %v", err), http.StatusServiceUnavailable)
		return
	}
	strDadosCotacao := string(dadosCotacao)
	fmt.Println(strDadosCotacao)
}

func main() {
	fmt.Println("Iniciando servidor HTTP...")

	http.HandleFunc("/cotacao", handlerCotacaoDolar)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
