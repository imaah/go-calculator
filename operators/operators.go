package operators

import "fmt"

type Operation interface {
	Eval() *OperationResult
}

type ResultType uint8

var OpResultNumber ResultType = 0
var OpResultString ResultType = 1

type OperationResult struct {
	resType  ResultType
	numValue float64
	strValue string
}

func (r OperationResult) IsNumber() bool {
	if r.resType == OpResultNumber {
		return true
	}
	return false
}

func (r OperationResult) IsString() bool {
	if r.resType == OpResultString {
		return true
	}
	return false
}

func (r OperationResult) GetNumber() float64 {
	return r.numValue
}

func (r OperationResult) GetString() string {
	if r.IsNumber() {
		return fmt.Sprintf("%f", r.GetNumber())
	}
	return r.strValue
}

func NewStringResult(value string) *OperationResult {
	return &OperationResult{
		resType:  OpResultString,
		strValue: value,
	}
}

func NewNumberResult(value float64) *OperationResult {
	return &OperationResult{
		resType:  OpResultNumber,
		numValue: value,
	}
}
