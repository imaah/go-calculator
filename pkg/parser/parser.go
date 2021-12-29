package parser

import (
	"emorisse.fr/go-calculator/pkg/operation"
	"emorisse.fr/go-calculator/pkg/operation/binary"
	"emorisse.fr/go-calculator/pkg/parser/builder"
	"emorisse.fr/go-calculator/pkg/parser/preprocessor"
)

func Parse(str string) (operation.Operation, error) {
	var m = make(map[uint]string)
	var n, _ = preprocessor.ProcessParenthesis(str, m)

	for _, token := range binary.OperatorPriority {
		var i = 0
		for i < len(m) {
			n, _ = preprocessor.ProcessToken(m[uint(i)], token, uint(i), n, m)
			i += 1
		}
	}

	return builder.BuildOperator(m, 0)
}
