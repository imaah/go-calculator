package operation_test

import (
	"github.com/imaah/go-calculator/pkg/operation"
	"math"
	"testing"
)

func TestNew_Function_WrongFunction(t *testing.T) {
	_, err := operation.NewFunction2("a", operation.NewNumber(1))

	if err == nil {
		t.Logf("Shouldn't be nil but go %s", err)
		t.Fail()
	}

}

func TestOpFunction_Eval(t *testing.T) {
	testFunction("cos", 0, math.Cos(0), t)
	testFunction("sin", 0, math.Sin(0), t)
	testFunction("tan", 0, math.Tan(0), t)
}

func TestNewUsingTempFunc(t *testing.T) {
	testFunc(wrap(func(val float64) float64 { return double(double(val)) }), 4, 16, t)
	testFunc(wrap(func(val float64) float64 { return val / 4 }), 4, 1, t)
	testFunc(wrap(func(val float64) float64 { return val }), 4, 4, t)
	testFunc(wrap(double), 4, 8, t)
}

func TestRegisterFunction_Function_Duplicate(t *testing.T) {
	err := operation.RegisterFunction("sin", wrap(double))

	if err == nil {
		t.Logf("Shouldn't be nil")
		t.Fail()
	}
}

func TestRegisterFunction(t *testing.T) {
	err := operation.RegisterFunction("double", wrap(double))

	if err != nil {
		t.Logf("Should be nil but got %s", err)
		t.Fail()
	}

	testFunction("double", 4, 8, t)
}

func wrap(fn func(float64) float64) operation.OpFunc {
	return operation.OpFunc{
		Func: func(o []operation.Operation) float64 {
			return fn(o[0].Eval())
		},
		NbArgs: 1,
	}
}

func double(val float64) float64 {
	return val * 2
}

func testFunc(function operation.OpFunc, val, expected float64, t *testing.T) {
	bin := operation.NewFunctionUsingTempFunc("temp", function, operation.NewNumber(val))

	if bin.Eval() != expected {
		t.Logf("Should be %f but got %f", expected, bin.Eval())
	}
}

func testFunction(function string, val, expected float64, t *testing.T) {
	bin, err := operation.NewFunction2(function, operation.NewNumber(val))

	if err != nil {
		t.Logf("Should be nil but got %s", err)
		t.Fail()
	}

	if bin.Eval() != expected {
		t.Logf("Should be %f but got %f", expected, bin.Eval())
	}
}
