package parser

import (
	"github.com/imaah/go-calculator/pkg/operation"
)

func ParseV2(str string) (operation.Operation, error) {
	tokens, err := Tokenize(str)

	if err != nil {
		return nil, err
	}

	pf, err := TokensToPostfix(tokens)

	if err != nil {
		return nil, err
	}

	op, err := Evaluate(pf)

	if err != nil {
		return nil, err
	}

	return op, nil
}
