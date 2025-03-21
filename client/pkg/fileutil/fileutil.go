package fileutil

import (
	"fmt"
	"os"
)

func WriteFileString(pathFileName string, content string) int {
	// Abrindo o arquivo com permissão de escrita e criação, se necessário
	// f, err := os.OpenFile(pathFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	f, err := os.OpenFile(pathFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Escrevendo uma string no arquivo e pegando o tamanho
	fileSize, err := f.WriteString(content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Escrita realizada com sucesso no arquivo", f.Name())
	return fileSize
}
