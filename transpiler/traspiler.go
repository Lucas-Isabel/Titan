package transpiler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Transpile(titanFilePath string) (string, error) {

	file, err := os.Open(titanFilePath)
	if err != nil {
		return "", fmt.Errorf("erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	var result strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		typechecked := false
		if strings.Contains(line, "local") {
			typechecked = true
			line = strings.ReplaceAll(line, ": ", "")
			line = strings.ReplaceAll(line, ":", "")
			line = strings.ReplaceAll(line, "string", "")
			line = strings.ReplaceAll(line, "int", "")
			line = strings.ReplaceAll(line, "float", "")
			line = strings.ReplaceAll(line, "bool", "")
			line = strings.ReplaceAll(line, "table", "")
		}
		if strings.Contains(line, "function") {
			typechecked = true
			line = strings.ReplaceAll(line, ": ", "")
			line = strings.ReplaceAll(line, ":", "")
			line = strings.ReplaceAll(line, "string", "")
			line = strings.ReplaceAll(line, "int", "")
			line = strings.ReplaceAll(line, "float", "")
			line = strings.ReplaceAll(line, "bool", "")
			line = strings.ReplaceAll(line, "table", "")
		}
		result.WriteString(line)
		if typechecked {
			result.WriteString(" -- removido a tipagem estática")
		}
		result.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("erro ao ler o arquivo: %v", err)
	}

	return result.String(), nil
}
