package operation

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var functions = make(map[string]OpFunc)

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
	_ = RegisterFunction("log10", math.Log10)
}

type OpFunc func(float64) float64

type OpFunction struct {
	FunctionName string
	Function     OpFunc
	Value        Operation
}

func (f OpFunction) Eval() float64 {
	return f.Function(f.Value.Eval())
}

func (f OpFunction) String() string {
	return f.FunctionName + "(" + f.Value.String() + ")"
}

// New Creates a new function operator
func NewFunction(functionName string, value Operation) (Operation, error) {
	if value == nil {
		return nil, errors.New("ArgumentIsNil")
	}
	lowFunName := strings.ToLower(functionName)
	if function, contains := functions[lowFunName]; contains {
		return &OpFunction{
			FunctionName: lowFunName,
			Function:     function,
			Value:        value,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("invalid function: %s", functionName))
}

func NewFunctionUsingTempFunc(function OpFunc, value Operation) Operation {
	return &OpFunction{
		FunctionName: "temp",
		Function:     function,
		Value:        value,
	}
}

// RegisterFunction Allows to register a new function for the function operator
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
