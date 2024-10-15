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
	Operation
	Left, Right Operation
	Symbol      rune
}

func (b *OpBinary) Eval() *Result {
	switch b.Symbol {
	case '+':
		return b.add(b.Right.Eval())
	case '-':
		return b.subtract(b.Right.Eval())
	case '/':
		return b.divide(b.Right.Eval())
	case '*':
		return b.multiply(b.Right.Eval())
	case '^':
		return b.pow(b.Right.Eval())
	case '%':
		return b.modulo(b.Right.Eval())
	}
	return nil
}

func (b OpBinary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.Left.String(), b.Symbol, b.Right.String())
}

func (b *OpBinary) multiply(right *Result) *Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() * right.GetNumber()
		return NewNumberResult(result)
	}

	return NewStringResult(left.GetString() + " * " + right.GetString())
}

func (b *OpBinary) divide(right *Result) *Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() / right.GetNumber()
		return NewNumberResult(result)
	}

	return NewStringResult(left.GetString() + " / " + right.GetString())
}

func (b *OpBinary) add(right *Result) *Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() + right.GetNumber()
		return NewNumberResult(result)
	}

	return NewStringResult(left.GetString() + " + " + right.GetString())
}

func (b *OpBinary) subtract(right *Result) *Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() - right.GetNumber()
		return NewNumberResult(result)
	}

	return NewStringResult(left.GetString() + " - " + right.GetString())
}

func (b *OpBinary) pow(right *Result) *Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = math.Pow(left.GetNumber(), right.GetNumber())
		return NewNumberResult(result)
	}

	return NewStringResult(left.GetString() + " ^ " + right.GetString())
}

func (b *OpBinary) modulo(right *Result) *Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = math.Mod(left.GetNumber(), right.GetNumber())
		return NewNumberResult(result)
	}

	return NewStringResult(left.GetString() + " % " + right.GetString())
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
		return &operator, nil
	}
	return nil, errors.New(fmt.Sprintf("invalid operator: %c", symbol))
}
