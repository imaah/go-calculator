package unary

import (
	"emorisse.fr/calcul/operators"
	"emorisse.fr/calcul/utils"
	"errors"
)

//KnownSymbols all symbols that can be used for a binary operator
var KnownSymbols = []rune{'-', '+'}

type opUnary struct {
	operators.Operation
	Right  operators.Operation
	Symbol rune
}

func (b *opUnary) Eval() float64 {
	switch b.Symbol {
	case '+':
		return b.Right.Eval()
	case '-':
		return -b.Right.Eval()
	}
	return 0
}

//New Creates a new unary operator
func New(symbol rune, right operators.Operation) (operators.Operation, error) {
	if right == nil {
		return nil, errors.New("ArgumentIsNil")
	}

	if utils.RuneArrayContains(KnownSymbols, symbol) {
		operator := opUnary{
			Right:  right,
			Symbol: symbol,
		}
		return &operator, nil
	}
	return nil, errors.New("InvalidOperator")
}
