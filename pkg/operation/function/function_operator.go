package function

import (
	"emorisse.fr/go-calculator/pkg/operation"
	"errors"
	"math"
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
}

type OpFunc func(float64) float64

type OpFunction struct {
	operation.Operation
	FunctionName string
	Function     OpFunc
	Value        operation.Operation
}

func (f OpFunction) Eval() *operation.Result {
	var innerRes = f.Value.Eval()

	if innerRes.IsNumber() {
		var res = f.Function(innerRes.GetNumber())
		return operation.NewNumberResult(res)
	}

	return operation.NewStringResult(f.FunctionName + "(" + innerRes.GetString() + ")")
}

func (f OpFunction) String() string {
	return f.FunctionName + "(" + f.Value.String() + ")"
}

//New Creates a new function operator
func New(functionName string, value operation.Operation) (operation.Operation, error) {
	if value == nil {
		return nil, errors.New("ArgumentIsNil")
	}
	if function, contains := functions[functionName]; contains {
		return &OpFunction{
			FunctionName: functionName,
			Function:     function,
			Value:        value,
		}, nil
	}
	return nil, errors.New("InvalidFunctionName")
}

func NewUsingTempFunc(function OpFunc, value operation.Operation) operation.Operation {
	return &OpFunction{
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
