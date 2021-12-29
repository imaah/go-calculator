package unary

import (
	"emorisse.fr/go-calculator/pkg/operation"
	"emorisse.fr/go-calculator/pkg/utils"
	"errors"
)

//KnownSymbols all symbols that can be used for a binary operator
var KnownSymbols = []rune{'-', '+'}

type OpUnary struct {
	operation.Operation
	Right  operation.Operation
	Symbol rune
}

func (b *OpUnary) Eval() *operation.Result {
	switch b.Symbol {
	case '+':
		return b.Right.Eval()
	case '-':
		return invert(b.Right.Eval())
	}
	return nil
}

func (b OpUnary) String() string {
	return string(b.Symbol) + b.Right.String()
}

func invert(right *operation.Result) *operation.Result {
	if right.IsNumber() {
		return operation.NewNumberResult(-right.GetNumber())
	}

	return operation.NewStringResult("-" + right.GetString())
}

//New Creates a new unary operator
func New(symbol rune, right operation.Operation) (operation.Operation, error) {
	if right == nil {
		return nil, errors.New("ArgumentIsNil (unary)")
	}

	if utils.RuneArrayContains(KnownSymbols, symbol) {
		operator := OpUnary{
			Right:  right,
			Symbol: symbol,
		}
		return &operator, nil
	}
	return nil, errors.New("InvalidOperator")
}
