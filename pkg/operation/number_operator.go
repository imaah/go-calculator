package operation

import (
	"fmt"
)

type OpNumber struct {
	Value float64
}

func (n OpNumber) Eval() float64 {
	return n.Value
}

func (n OpNumber) String() string {
	return fmt.Sprintf("%f", n.Value)
}

// New Creates a new number operator
func NewNumber(value float64) Operation {
	return OpNumber{
		Value: value,
	}
}
