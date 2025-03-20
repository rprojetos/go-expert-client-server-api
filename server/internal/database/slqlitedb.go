package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

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

// createTableCotacoes cria a tabela "cotacoes" e aguarda até 200ms para garantir que esteja disponível.
func createTableIfNotExistsCotacoes(db *sql.DB) error {
	// Criando a tabela se não existir
	createTableQuery := `
        CREATE TABLE IF NOT EXISTS cotacoes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        code TEXT,
        codein TEXT,
        name TEXT,
        high TEXT,
        low TEXT,
        var_bid TEXT,
        pct_change TEXT,
        bid TEXT,
        ask TEXT,
        timestamp TEXT,
        create_date TEXT,
        data_consulta DATETIME DEFAULT CURRENT_TIMESTAMP
    )`
	
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("erro ao criar tabela: %v", err)
	}
	fmt.Println("Comando de criação da tabela enviado.")

	// Aguarda até que a tabela esteja disponível (máx. 200ms)
	maxWait := 200 * time.Millisecond
	interval := 50 * time.Millisecond
	startTime := time.Now()

	for {
		if tableExists(db, "cotacoes") {
			fmt.Println("Tabela criada com sucesso!")
			return nil // Sucesso
		}

		// Se ultrapassar o tempo máximo, retorna erro
		if time.Since(startTime) > maxWait {
			return fmt.Errorf("tempo limite atingido, a tabela pode não ter sido criada corretamente")
		}

		time.Sleep(interval) // Aguarda 50ms antes de tentar novamente
	}
}

// tableExists verifica se uma tabela existe no banco de dados SQLite
func tableExists(db *sql.DB, tableName string) bool {
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name=?"
	var name string
	err := db.QueryRow(query, tableName).Scan(&name)

	return err == nil
}

func parseJsonCotacao(dadosCotacao []byte, resultado *Resultado) (*Resultado, error){
	if err := json.Unmarshal(dadosCotacao, &resultado); err != nil {
		return nil, fmt.Errorf("erro ao processar dados para salvar: %v", err)
	}
	return resultado, nil
}

func SaveDadosCotacao(dadosCotacao []byte) error {
	// TODO deletar o startTime e o endTime
	startTime := time.Now()
	// Cria um contexto com timeout de 10 milissegundos
	// TODO analisar tempo de 10 mS
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel() // Garante que os recursos sejam liberados quando a função retornar
	// Abre conexão com o banco SQLite
	db, err := sql.Open("sqlite", "data/finance.db")
	if err != nil {
		return fmt.Errorf("erro ao abrir banco de dados: %v", err)
	}
	defer db.Close()

	if !tableExists(db, "cotacoes") {
		err := createTableIfNotExistsCotacoes(db)
		if err != nil {
			return fmt.Errorf("erro ao preparar statement: %v", err)
		}
	}

	// Prepara a statement usando o contexto com timeout
	stmt, err := db.PrepareContext(ctx, `INSERT INTO cotacoes (
        code, codein, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("erro ao preparar statement: %v", err)
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
		return fmt.Errorf("erro ao inserir dados: %v", err)
	}

	endTime := time.Now()

	elapsed := endTime.Sub(startTime)
	fmt.Println("Tempo de execução:", elapsed)

	return nil
}
