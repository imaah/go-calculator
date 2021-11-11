package number

import "emorisse.fr/calcul/operators"

type opNumber struct {
	operators.Operation
	Value float64
}

func (n opNumber) Eval() float64 {
	return n.Value
}

//New Creates a new number operator
func New(value float64) operators.Operation {
	return opNumber{
		Value: value,
	}
}
