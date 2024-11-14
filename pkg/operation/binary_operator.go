package operation

import (
	"errors"
	"fmt"
	"math"
	"slices"
)

// OperatorPriority all symbols that can be used for a binary operator ordered by priority
var OperatorPriority = []rune{'^', '*', '/', '%', '-', '+'}

type OpBinary struct {
	Left, Right Operation
	Symbol      rune
}

func (b OpBinary) Eval() float64 {
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

func (b OpBinary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.Left.String(), b.Symbol, b.Right.String())
}

// New Creates a new binary operator
func NewBinary(symbol rune, left, right Operation) (Operation, error) {
	if left == nil || right == nil {
		return nil, errors.New("ArgumentIsNil (binary)")
	}

	if slices.Contains(OperatorPriority, symbol) {
		operator := OpBinary{
			Right:  right,
			Left:   left,
			Symbol: symbol,
		}
		return operator, nil
	}
	return nil, errors.New(fmt.Sprintf("invalid binary operator: %c", symbol))
}
