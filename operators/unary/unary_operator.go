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

func (b *opUnary) Eval() *operators.OperationResult {
	switch b.Symbol {
	case '+':
		return b.Right.Eval()
	case '-':
		return invert(b.Right.Eval())
	}
	return nil
}

func invert(right *operators.OperationResult) *operators.OperationResult {
	if right.IsNumber() {
		return operators.NewNumberResult(-right.GetNumber())
	}

	return operators.NewStringResult("-" + right.GetString())
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
