package main

import (
	"fmt"
	"os"
	"titan/interpreter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: titan <arquivo.titan>")
		return
	}

	filePath := os.Args[1]
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo: %v\n", err)
		return
	}

	interpreter.Run(string(content))
}
