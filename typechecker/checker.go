package typechecker

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type TitanVariable struct {
	Name  string
	Type  string
	Value string
}

func ExtractVariables(filePath string) ([]TitanVariable, error) {
	reg := regexp.MustCompile(`^\s*local\s+(\w+)\s*:\s*(\w+)\s*=\s*(.+)$`)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	defer file.Close()

	var variables []TitanVariable
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		match := reg.FindStringSubmatch(line)
		if len(match) == 4 {
			variable := TitanVariable{
				Name:  strings.TrimSpace(match[1]),
				Type:  strings.TrimSpace(match[2]),
				Value: strings.TrimSpace(match[3]),
			}
			variables = append(variables, variable)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("erro ao ler linhas do arquivo: %w", err)
	}

	return variables, nil
}

// Verifica se o valor declarado bate com o tipo especificado
func (v TitanVariable) Check() error {
	switch v.Type {
	case "string":
		if !(strings.HasPrefix(v.Value, "\"") && strings.HasSuffix(v.Value, "\"")) {
			return fmt.Errorf("variável '%s': valor '%s' não é uma string válida", v.Name, v.Value)
		}
	case "int":
		if _, err := strconv.Atoi(v.Value); err != nil {
			return fmt.Errorf("variável '%s': valor '%s' não é um int válido", v.Name, v.Value)
		}
	case "float":
		if _, err := strconv.ParseFloat(v.Value, 64); err != nil {
			return fmt.Errorf("variável '%s': valor '%s' não é um float válido", v.Name, v.Value)
		}
	case "bool":
		lower := strings.ToLower(v.Value)
		if lower != "true" && lower != "false" {
			return fmt.Errorf("variável '%s': valor '%s' não é um bool válido", v.Name, v.Value)
		}
	default:
		return errors.New("tipo desconhecido: " + v.Type)
	}
	return nil
}
