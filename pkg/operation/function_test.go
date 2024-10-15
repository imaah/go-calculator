package operation_test

import (
	"github.com/imaah/go-calculator/pkg/operation"
	"math"
	"testing"
)

func TestNew_Function_WrongFunction(t *testing.T) {
	var _, err = operation.NewFunction("a", operation.NewNumber(1))

	if err == nil {
		t.Logf("Shouldn't be nil but go %s", err)
		t.Fail()
	}

}

func TestNew_Function_NilArgument(t *testing.T) {
	var _, err = operation.NewFunction("cos", nil)

	if err == nil {
		t.Logf("Shouldn't be nil")
		t.Fail()
	}
}

func TestOpFunction_Eval(t *testing.T) {
	testFunction("cos", 0, math.Cos(0), t)
	testFunction("sin", 0, math.Sin(0), t)
	testFunction("tan", 0, math.Tan(0), t)
}

func TestNewUsingTempFunc(t *testing.T) {
	testFunc(func(val float64) float64 { return double(double(val)) }, 4, 16, t)
	testFunc(func(val float64) float64 { return val / 4 }, 4, 1, t)
	testFunc(func(val float64) float64 { return val }, 4, 4, t)
	testFunc(double, 4, 8, t)
}

func TestRegisterFunction_Function_NilArgument(t *testing.T) {
	var err = operation.RegisterFunction("double", nil)

	if err == nil {
		t.Logf("Shouldn't be nil")
		t.Fail()
	}
}

func TestRegisterFunction_Function_Duplicate(t *testing.T) {
	var err = operation.RegisterFunction("sin", double)

	if err == nil {
		t.Logf("Shouldn't be nil")
		t.Fail()
	}
}

func TestRegisterFunction(t *testing.T) {
	var err = operation.RegisterFunction("double", double)

	if err != nil {
		t.Logf("Should be nil but got %s", err)
		t.Fail()
	}

	testFunction("double", 4, 8, t)
}

func double(val float64) float64 {
	return val * 2
}

func testFunc(function operation.OpFunc, val, expected float64, t *testing.T) {
	var bin = operation.NewFunctionUsingTempFunc(function, operation.NewNumber(val))

	if bin.Eval().GetNumber() != expected {
		t.Logf("Should be %f but got %f", expected, bin.Eval().GetNumber())
	}
}

func testFunction(function string, val, expected float64, t *testing.T) {
	var bin, err = operation.NewFunction(function, operation.NewNumber(val))

	if err != nil {
		t.Logf("Should be nil but got %s", err)
		t.Fail()
	}

	if bin.Eval().GetNumber() != expected {
		t.Logf("Should be %f but got %f", expected, bin.Eval().GetNumber())
	}
}
