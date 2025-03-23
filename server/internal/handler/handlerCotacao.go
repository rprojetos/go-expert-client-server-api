package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rprojetos/go-expert/internal/config"
	"github.com/rprojetos/go-expert/internal/database"
	"github.com/rprojetos/go-expert/pkg/cotacaoapi"
)

func HandlerCotacaoDolar(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	// Verifica se o método é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido. Use POST.", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os dados da cotação
	dadosCotacao, err := cotacaoapi.ObterCotacaoDolar()
	if err != nil {
		http.Error(w, fmt.Sprintf("Falha ao obter cotação: %v", err), http.StatusServiceUnavailable)
		return
	}

	if err := database.SaveDadosCotacao(dadosCotacao); err != nil {
		if err == context.DeadlineExceeded {
			cfg, err := config.LoadConfig()
			if err != nil {
				fmt.Printf("Erro ao ler configuração de timeout do db %v\n", err)
			}
			timeDbSqlite := cfg.Context.Timeout.TimeDbSqlite
			log.Printf("ERRO: Timeout na realização de persistência de dados - limite de %dms excedido", timeDbSqlite)

		}
		// Loga o erro, mas continua o processamento para retornar o bid
		log.Printf("Erro ao salvar dados: %v\n", err)
	}

	// Estrutura para parsear o JSON
	var resultado struct {
		USDBRL struct {
			Bid string `json:"bid"`
		} `json:"USDBRL"`
	}

	// Faz o parse do JSON
	if err := json.Unmarshal(dadosCotacao, &resultado); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao processar dados: %v", err), http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho da resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Cria e envia a resposta contendo apenas o valor bid
	resposta := map[string]string{"bid": resultado.USDBRL.Bid}
	json.NewEncoder(w).Encode(resposta)

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	log.Println("Tempo total de execução/resposta:", elapsed)
}
