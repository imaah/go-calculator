package operation

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var functions = make(map[string]OpFunc)

func init() {
	_ = RegisterFunction("cos", wrap1(math.Cos))
	_ = RegisterFunction("sin", wrap1(math.Sin))
	_ = RegisterFunction("tan", wrap1(math.Tan))
	_ = RegisterFunction("sqrt", wrap1(math.Sqrt))
	_ = RegisterFunction("atan", wrap1(math.Atan))
	_ = RegisterFunction("acos", wrap1(math.Acos))
	_ = RegisterFunction("asin", wrap1(math.Asin))
	_ = RegisterFunction("log", wrap1(math.Log))
	_ = RegisterFunction("ceil", wrap1(math.Ceil))
	_ = RegisterFunction("floor", wrap1(math.Floor))
	_ = RegisterFunction("round", wrap1(math.Round))
	_ = RegisterFunction("exp", wrap1(math.Exp))
	_ = RegisterFunction("log10", wrap1(math.Log10))
	_ = RegisterFunction("abs", wrap1(math.Abs))
	_ = RegisterFunction("mod", wrap2(math.Mod))
	_ = RegisterFunction("atan2", wrap2(math.Atan2))
	_ = RegisterFunction("min", wrapV(min))
	_ = RegisterFunction("max", wrapV(max))
}

func min(x float64, y ...float64) float64 {
	m := x

	for _, v := range y {
		if v < m {
			m = v
		}
	}
	return m
}

func max(x float64, y ...float64) float64 {
	m := x

	for _, v := range y {
		if v > m {
			m = v
		}
	}
	return m
}

func wrap1(fn func(float64) float64) OpFunc {
	return OpFunc{
		NbArgs: 1,
		Func: func(f []Operation) float64 {
			return fn(f[0].Eval())
		},
	}
}

func wrap2(fn func(float64, float64) float64) OpFunc {
	return OpFunc{
		NbArgs: 2,
		Func: func(f []Operation) float64 {
			return fn(f[0].Eval(), f[1].Eval())
		},
	}
}

func wrapV(fn func(float64, ...float64) float64) OpFunc {
	return OpFunc{
		NbArgs:         0,
		VariableNbArgs: true,
		Func: func(o []Operation) float64 {
			vals := make([]float64, 0, len(o))
			for _, v := range o {
				vals = append(vals, v.Eval())
			}
			return fn(vals[0], vals[1:]...)
		},
	}
}

type OpFunc struct {
	Func           func([]Operation) float64
	NbArgs         int
	VariableNbArgs bool
}

type OpFunction struct {
	FunctionName string
	Function     OpFunc
	Values       []Operation
}

func (f OpFunction) Eval() float64 {
	return f.Function.Func(f.Values)
}

func (f OpFunction) String() string {
	vals := make([]string, 0, len(f.Values))

	for _, op := range f.Values {
		vals = append(vals, op.String())
	}

	return fmt.Sprintf("%s(%s)", f.FunctionName, strings.Join(vals, ","))
}

func NewFunction2(name string, values ...Operation) (Operation, error) {
	lowFunName := strings.ToLower(name)
	function, contains := functions[lowFunName]

	if !contains {
		return nil, errors.New(fmt.Sprintf("invalid function: %s", name))
	}

	if !function.VariableNbArgs && len(values) != function.NbArgs {
		return nil, fmt.Errorf("invalid number of arguments for function %s, got %d expected %d", name, len(values), function.NbArgs)
	}

	return OpFunction{
		FunctionName: lowFunName,
		Function:     function,
		Values:       values,
	}, nil
}

func NewFunctionUsingTempFunc(name string, function OpFunc, value Operation) Operation {
	return OpFunction{
		FunctionName: name,
		Function:     function,
		Values:       []Operation{value},
	}
}

// RegisterFunction Allows to register a new function for the function operator
func RegisterFunction(name string, function OpFunc) error {
	if _, contains := functions[name]; contains {
		return errors.New("AlreadyRegisteredFunction")
	}
	functions[name] = function
	return nil
}
