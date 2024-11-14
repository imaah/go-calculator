package operation

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var constants = make(map[string]float64)

func init() {
	_ = RegisterConstant("pi", math.Pi)
	_ = RegisterConstant("e", math.E)
	_ = RegisterConstant("exp", math.E)
	_ = RegisterConstant("phi", math.Phi)
	_ = RegisterConstant("nan", math.NaN())
}

type OpConstant struct {
	Name  string
	Value float64
}

func (f OpConstant) Eval() float64 {
	return f.Value
}

func (f OpConstant) String() string {
	return f.Name
}

func NewConstant(name string) (Operation, error) {
	lowConstName := strings.ToLower(name)
	constant, contains := constants[lowConstName]

	if !contains {
		return nil, errors.New(fmt.Sprintf("invalid constant: %s", name))
	}

	return OpConstant{
		Name:  name,
		Value: constant,
	}, nil
}

func RegisterConstant(name string, constant float64) error {
	if _, contains := constants[name]; contains {
		return errors.New("AlreadyRegisteredFunction")
	}
	constants[name] = constant
	return nil
}
