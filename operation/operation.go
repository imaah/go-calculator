package operation

import "fmt"

type Operation interface {
	Eval() *Result
	String() string
}

type ResultType uint8

var OpResultNumber ResultType = 0
var OpResultString ResultType = 1

type Result struct {
	resType  ResultType
	numValue float64
	strValue string
}

func (r Result) IsNumber() bool {
	if r.resType == OpResultNumber {
		return true
	}
	return false
}

func (r Result) IsString() bool {
	if r.resType == OpResultString {
		return true
	}
	return false
}

func (r Result) GetNumber() float64 {
	return r.numValue
}

func (r Result) GetString() string {
	if r.IsNumber() {
		return fmt.Sprintf("%f", r.GetNumber())
	}
	return r.strValue
}

func NewStringResult(value string) *Result {
	return &Result{
		resType:  OpResultString,
		strValue: value,
	}
}

func NewNumberResult(value float64) *Result {
	return &Result{
		resType:  OpResultNumber,
		numValue: value,
	}
}
