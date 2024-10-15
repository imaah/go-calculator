package operation

import (
	"fmt"
)

type OpNumber struct {
	Operation
	value float64
}

func (n OpNumber) Eval() *Result {
	return NewNumberResult(n.value)
}

func (n OpNumber) String() string {
	return fmt.Sprintf("%f", n.value)
}

// New Creates a new number operator
func NewNumber(value float64) Operation {
	return OpNumber{
		value: value,
	}
}
