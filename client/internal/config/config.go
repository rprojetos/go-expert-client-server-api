package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Estrutura para armazenar as configurações do arquivo YAML
type Config struct {
	Cotacao struct {
		Url string `yaml:"url"`
		PathFileName string `yaml:"pathFileName"`
	} `yaml:"cotacao"`
	Context struct {
		Timeout struct {
			TimeResponseApi int `yaml:"timeResponseApi"`
		} `yaml:"timeout"`
	} `yaml:"context"`
}

// LoadConfig lê o arquivo config.yaml e retorna a configuração
func LoadConfig() (Config, error) {
	var cfg Config
	filePath := "internal/config/config.yaml" // Caminho do arquivo

	// Lê o conteúdo do arquivo YAML
	data, err := os.ReadFile(filePath)
	if err != nil {
		return cfg, fmt.Errorf("erro ao carregar configuração: %w", err)
	}

	// Faz o parse do YAML
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("erro ao processar configuração: %w", err)
	}

	return cfg, nil
}
