package bootstrap

import (
	"database/sql"
	"log"
	"time"
)

// createTableCotacoes cria a tabela "cotacoes" e aguarda até 200ms para garantir que esteja disponível.
func createTableIfNotExistsCotacoes() error {
	db, err := sql.Open("sqlite", "data/finance.db")
	if err != nil {
		return err
	}
	defer db.Close()

	if tableExists(db, "cotacoes") {
		log.Println("Ok!")
		return nil
	}

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

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return err
	}
	log.Println("Criando a tabela cotacoes no banco de dados finance.db.")

	// Aguarda até que a tabela esteja disponível (máx. 200ms)
	maxWait := 200 * time.Millisecond
	interval := 50 * time.Millisecond
	startTime := time.Now()

	for {
		if tableExists(db, "cotacoes") {
			log.Println("Tabela criada com sucesso!")
			return nil
		}

		// Se ultrapassar o tempo máximo, retorna erro
		if time.Since(startTime) > maxWait {
			return err
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

func ConfigSystem() error {
	log.Println("Verificando e preparando o sistema...")
	err := createTableIfNotExistsCotacoes()
	if err != nil {
		return err
	}
	return nil
}
