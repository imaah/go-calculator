package binary

import (
	"emorisse.fr/calcul/operators"
	"emorisse.fr/calcul/utils"
	"errors"
	"math"
)

//KnownSymbols all symbols that can be used for a binary operator
var KnownSymbols = []rune{'-', '+', '*', '/', '^', '%'}

var OperatorPriority = [][]rune{
	{'*', '/', '^', '%'},
	{'-', '+'},
}

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
	case '^':
		return math.Pow(b.Left.Eval(), b.Right.Eval())
	case '%':
		return math.Mod(b.Left.Eval(), b.Right.Eval())
	}
	return 0
}

//New Creates a new binary operator
func New(symbol rune, left, right operators.Operation) (operators.Operation, error) {
	if left == nil || right == nil {
		return nil, errors.New("ArgumentIsNil")
	}

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
