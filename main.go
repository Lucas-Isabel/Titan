package main

import (
	"fmt"
	"os"
	"titan/runtime"
	"titan/transpiler"
	"titan/typechecker"
)

func SaveToFile(outputPath, content string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo: %v", err)
	}

	return nil
}

func main() {
	path := "D:/Titan/exemplos/error_type.tn.lua"
	output := "D:/Titan/output/hello_world.lua"

	check, err := typechecker.ExtractVariables(path)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Declarações de variáveis:", check)

	for _, variable := range check {
		if err := variable.Check(); err != nil {
			fmt.Println("Erro:", err)
			return
		}
	}

	result, err := transpiler.Transpile(path)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	if err := SaveToFile(output, result); err != nil {
		fmt.Println("Erro:", err)
		return
	}

	vmErr := runtime.Run(output)
	if vmErr != nil {
		fmt.Println("Erro:", vmErr)
		return
	}
}
