package fileutil

import (
	"log"
	"os"
)

func WriteFileString(pathFileName string, content string) (int, error) {
	// Abrindo o arquivo com permissão de escrita e criação, se necessário
	f, err := os.OpenFile(pathFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// Escrevendo uma string no arquivo e pegando o tamanho
	fileSize, err := f.WriteString(content)
	if err != nil {
		return 0, err
	}

	log.Printf("Escrita realizada com sucesso no arquivo %s", f.Name())
	return fileSize, nil
}
