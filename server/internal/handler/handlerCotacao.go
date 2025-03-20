package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rprojetos/go-expert/internal/database"
	"github.com/rprojetos/go-expert/pkg/cotacaoapi"
)

func HandlerCotacaoDolar(w http.ResponseWriter, r *http.Request) {
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
		// Loga o erro, mas continua o processamento para retornar o bid
		fmt.Printf("Erro ao salvar dados: %v\n", err)
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
}
