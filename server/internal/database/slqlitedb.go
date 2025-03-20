package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rprojetos/go-expert/internal/config"
	_ "modernc.org/sqlite"
)

// Estrutura para parsear o JSON completo
type Resultado struct {
	USDBRL struct {
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
	} `json:"USDBRL"`
}

func parseJsonCotacao(dadosCotacao []byte, resultado *Resultado) (*Resultado, error) {
	if err := json.Unmarshal(dadosCotacao, &resultado); err != nil {
		return nil, fmt.Errorf("erro ao processar dados para salvar: %v", err)
	}
	return resultado, nil
}

func SaveDadosCotacao(dadosCotacao []byte) error {
	startTime := time.Now()
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	timeDbSqlite := cfg.Context.Timeout.TimeDbSqlite
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeDbSqlite)*time.Millisecond)
	defer cancel() // Garante que os recursos sejam liberados quando a função retornar
	// Abre conexão com o banco SQLite
	db, err := sql.Open("sqlite", "data/finance.db")
	if err != nil {
		return fmt.Errorf("erro ao abrir banco de dados: %v", err)
	}
	defer db.Close()

	// Prepara a statement usando o contexto com timeout
	stmt, err := db.PrepareContext(ctx, `INSERT INTO cotacoes (
        code, codein, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		fmt.Println("Erro ao preparar statement")
		return err
	}
	defer stmt.Close()

	var resultado Resultado
	// Faz o parse do JSON
	parseJsonCotacao(dadosCotacao, &resultado)

	// Executa a statement usando o contexto com timeout
	_, err = stmt.ExecContext(ctx,
		resultado.USDBRL.Code,
		resultado.USDBRL.Codein,
		resultado.USDBRL.Name,
		resultado.USDBRL.High,
		resultado.USDBRL.Low,
		resultado.USDBRL.VarBid,
		resultado.USDBRL.PctChange,
		resultado.USDBRL.Bid,
		resultado.USDBRL.Ask,
		resultado.USDBRL.Timestamp,
		resultado.USDBRL.CreateDate,
	)
	if err != nil {
		fmt.Println("Erro ao inserir dados.")
		return err
	}

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Println("Tempo de persistência no db:", elapsed)

	return nil
}
