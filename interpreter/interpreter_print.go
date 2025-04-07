package interpreter

import (
	"strings"
	"titan/output"
)

func processPrint(line string, state *InterpreterState) {
	params := strings.TrimSuffix(strings.TrimPrefix(line, "print("), ")")

	parts := strings.Split(params, ",")
	var values []interface{}

	for _, part := range parts {
		part = strings.TrimSpace(part)
		val := evaluateExpression(part, state)
		values = append(values, val)
	}

	output.Print(values...)
}
