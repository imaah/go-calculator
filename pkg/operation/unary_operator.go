package operation

import (
	"errors"
	"fmt"
)

type OpUnary struct {
	Right  Operation
	Symbol rune
}

func (b OpUnary) Eval() float64 {
	switch b.Symbol {
	case '-':
		return -b.Right.Eval()
	case '+':
		return b.Right.Eval()
	}
	return 0
}

func (b OpUnary) String() string {
	if b.Symbol == '-' {
		return fmt.Sprintf("-(%s)", b.Right.String())
	}
	return b.Right.String()
}

func NewUnary(symbol rune, right Operation) (Operation, error) {
	fmt.Println(right)
	if right == nil {
		return nil, errors.New("argument is nil")
	}

	if symbol != '-' && symbol != '+' {
		return nil, fmt.Errorf("invalid unary operator: %c", symbol)
	}

	return OpUnary{Symbol: symbol, Right: right}, nil
}
