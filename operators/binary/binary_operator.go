package binary

import (
	"emorisse.fr/calcul/operators"
	"emorisse.fr/calcul/utils"
	"errors"
)

//KnownSymbols all symbols that can be used for a binary operator
var KnownSymbols = []rune{'-', '+', '*', '/'}

type opBinary struct {
	operators.Operation
	Left, Right operators.Operation
	Symbol      rune
}

func (b *opBinary) Eval() float64 {
	switch b.Symbol {
	case '+':
		return b.Left.Eval() + b.Right.Eval()
	case '-':
		return b.Left.Eval() - b.Right.Eval()
	case '/':
		return b.Left.Eval() / b.Right.Eval()
	case '*':
		return b.Left.Eval() * b.Right.Eval()
	}
	return 0
}

//New Creates a new binary operator
func New(symbol rune, left, right operators.Operation) (operators.Operation, error) {
	if utils.RuneArrayContains(KnownSymbols, symbol) {
		operator := opBinary{
			Right:  right,
			Left:   left,
			Symbol: symbol,
		}
		return &operator, nil
	}
	return nil, errors.New("InvalidOperator")
}
