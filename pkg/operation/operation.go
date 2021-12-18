package operation

import "fmt"

var OpResultNumber ResultType = 0
var OpResultString ResultType = 1

type ResultType uint8

type Operation interface {
	Eval() *Result
	String() string
}

type Result struct {
	resType  ResultType
	numValue float64
	strValue string
}

func (r Result) IsNumber() bool {
	return r.resType == OpResultNumber
}

func (r Result) IsString() bool {
	return r.resType == OpResultString
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
