package function

import (
	"emorisse.fr/calcul/operators"
	"errors"
	"math"
)

type OpFunc func(float64) float64

var functions = make(map[string]OpFunc)

type opFunction struct {
	operators.Operation
	FunctionName string
	Function     OpFunc
	Value        operators.Operation
}

func (f opFunction) Eval() float64 {
	return f.Function(f.Value.Eval())
}

func init() {
	_ = RegisterFunction("cos", math.Cos)
	_ = RegisterFunction("sin", math.Sin)
	_ = RegisterFunction("tan", math.Tan)
	_ = RegisterFunction("sqrt", math.Sqrt)
	_ = RegisterFunction("atan", math.Atan)
	_ = RegisterFunction("acos", math.Acos)
	_ = RegisterFunction("asin", math.Asin)
	_ = RegisterFunction("log", math.Log)
	_ = RegisterFunction("ceil", math.Ceil)
	_ = RegisterFunction("floor", math.Floor)
	_ = RegisterFunction("round", math.Round)
	_ = RegisterFunction("exp", math.Exp)
}

//New Creates a new function operator
func New(functionName string, value operators.Operation) (operators.Operation, error) {
	if value == nil {
		return nil, errors.New("ArgumentIsNil")
	}
	if function, contains := functions[functionName]; contains {
		return &opFunction{
			FunctionName: functionName,
			Function:     function,
			Value:        value,
		}, nil
	}
	return nil, errors.New("InvalidFunctionName")
}

func NewUsingTempFunc(function OpFunc, value operators.Operation) operators.Operation {
	return &opFunction{
		FunctionName: "temp",
		Function:     function,
		Value:        value,
	}
}

//RegisterFunction Allows to register a new function for the function operator
func RegisterFunction(name string, function OpFunc) error {
	if function == nil {
		return errors.New("ArgumentIsNil")
	}
	if _, contains := functions[name]; contains {
		return errors.New("AlreadyRegisteredFunction")
	}
	functions[name] = function
	return nil
}
