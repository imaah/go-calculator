package parser

import (
	"fmt"
)

type Token interface {
	Pred() int
	Asso() byte
}

type Number float64

func (n Number) String() string {
	return fmt.Sprintf("Number(%.1f)", n)
}

func (n Number) Pred() int {
	return 0
}

func (n Number) Asso() byte {
	return 'L'
}

type Binary rune

func (b Binary) String() string {
	return fmt.Sprintf("Operator(%c)", rune(b))
}

func (b Binary) Pred() int {
	switch b {
	case '+', '-':
		return 1
	case '*', '/', '%':
		return 2
	case '^':
		return 3
	}
	return -1
}

func (b Binary) Asso() byte {
	if b == '^' {
		return 'R'
	}
	return 'L'
}

type Function string

func (f Function) String() string {
	return fmt.Sprintf("Function(%s)", string(f))
}

func (f Function) Pred() int {
	return -1
}

func (Function) Asso() byte {
	return 'L'
}

type LParen struct{}

func (LParen) String() string {
	return "LParen"
}

func (LParen) Pred() int {
	return -1
}

func (LParen) Asso() byte {
	return 'L'
}

type RParen struct{}

func (RParen) String() string {
	return "RParen"
}

func (RParen) Pred() int {
	return -1
}

func (RParen) Asso() byte {
	return 'L'
}
