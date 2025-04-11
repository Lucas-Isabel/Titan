package interpreter

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type TitanValue interface{}

type TitanTable struct {
	fields map[string]TitanValue
}

func NewTitanTable() *TitanTable {
	return &TitanTable{fields: make(map[string]TitanValue)}
}

func (t *TitanTable) Get(key string) TitanValue {
	return t.fields[key]
}

func (t *TitanTable) Set(key string, value TitanValue) {
	t.fields[key] = value
}

type InterpreterState struct {
	Variables     map[string]TitanValue
	VariableTypes map[string]string
	Count         int
}

func NewInterpreterState() *InterpreterState {
	return &InterpreterState{
		Variables:     make(map[string]TitanValue),
		VariableTypes: make(map[string]string),
	}
}

// var variables = make(map[string]TitanValue)
// var variableTypes = make(map[string]string)
// var state.Count int

func Run(code string) {
	state := NewInterpreterState()
	variables := state.Variables
	variableTypes := state.VariableTypes

	lines := strings.Split(code, "\n")
	for _, line := range lines {
		state.Count++
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "print(") {
			// Processa a linha de impressão
			processPrint(line, state)
			continue
		}

		if strings.HasPrefix(line, "local ") {
			line = strings.TrimPrefix(line, "local ")
			parts := strings.SplitN(line, "=", 2)
			left := strings.TrimSpace(parts[0])
			right := strings.TrimSpace(parts[1])

			varName, typeHint := parseVariableDeclaration(left)

			if typeHint == "" {
				message := fmt.Sprintf("Error: type hint is required in variable declaration")
				fmt.Println("Error: type hint is required in variable declaration")
				Error(2, message, state)
				continue
			}

			value := evaluateExpression(right, state)

			// Verifica se o tipo da variável corresponde ao tipo indicado
			if !checkType(typeHint, value) {
				message := fmt.Sprintf("Error: Type mismatch for variable %s. Expected %s, got %T", varName, typeHint, value)
				fmt.Printf("Error: Type mismatch for variable %s. Expected %s, got %T\n", varName, typeHint, value)
				Error(1, message, state)
				continue
			}

			variableTypes[varName] = typeHint
			variables[varName] = value
			continue
		}

		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			varName := strings.TrimSpace(parts[0])
			value := evaluateExpression(strings.TrimSpace(parts[1]), state)

			// Verifica se a variável tem um tipo definido
			if typeHint, ok := variableTypes[varName]; ok {
				if !checkType(typeHint, value) {
					message := fmt.Sprintf("Error: Type mismatch for variable %s. Expected %s, got %T", varName, typeHint, value)
					fmt.Printf("Error: Type mismatch for variable %s. Expected %s, got %T\n", varName, typeHint, value)
					Error(1, message, state)
					continue
				}
			} else {
				message := fmt.Sprintf("Error: Variable '%s' is not defined with a type", varName)
				fmt.Printf("Error: Variable '%s' is not defined with a type\n", varName)
				Error(3, message, state)
				continue
			}

			variables[varName] = value
		}
	}
}

func parseVariableDeclaration(declaration string) (string, string) {
	// Separa o nome da variável e o tipo
	if strings.Contains(declaration, ":") {
		parts := strings.Split(declaration, ":")
		varName := strings.TrimSpace(parts[0])
		typeHint := strings.TrimSpace(parts[1])
		return varName, typeHint
	}
	return declaration, ""
}

func evaluateExpression(expr string, state *InterpreterState) TitanValue {
	variables := state.Variables
	expr = strings.TrimSpace(expr)

	if val, err := strconv.Atoi(expr); err == nil {
		return val
	}

	if val, err := strconv.ParseFloat(expr, 64); err == nil {
		return val
	}

	if strings.HasPrefix(expr, "\"") && strings.HasSuffix(expr, "\"") {
		return strings.Trim(expr, "\"")
	}

	if strings.Contains(expr, "+") {
		parts := strings.SplitN(expr, "+", 2)
		left := evaluateExpression(parts[0], state)
		right := evaluateExpression(parts[1], state)

		switch l := left.(type) {
		case int:
			switch r := right.(type) {
			case int:
				return l + r
			case float64:
				return float64(l) + r
			}
		case float64:
			switch r := right.(type) {
			case int:
				return l + float64(r)
			case float64:
				return l + r
			}
		case string:
			return l + fmt.Sprint(right)
		}
	}

	if val, ok := variables[expr]; ok {
		return val
	}

	Error(4, fmt.Sprintf("Variable '%s' not found", expr), state)
	return nil
}

func checkType(typeHint string, value TitanValue) bool {
	switch typeHint {
	case "int":
		_, ok := value.(int)
		return ok
	case "float":
		_, ok1 := value.(float64)
		_, ok2 := value.(int)
		return ok1 || ok2
	case "string":
		_, ok := value.(string)
		return ok
	case "nil":
		return value == nil
	case "any":
		return true
	default:
		return true
	}
}

var typeErrors = map[int]string{
	1: "Invalid variable type",
	2: "Type mismatch",
	3: "Variable not defined",
	4: "Unsupported operation",
}

func Error(typeError int, message string, state *InterpreterState) {
	errorMessage, exists := typeErrors[typeError]
	if !exists {
		errorMessage = "Unknown error"
	}

	log.Fatalf("Error [%d] in line %d: %s, %s", typeError, state.Count, errorMessage, message)
}
