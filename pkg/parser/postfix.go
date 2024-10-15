package parser

import (
	"fmt"

	"github.com/imaah/go-calculator/pkg/operation"
	"github.com/imaah/go-calculator/pkg/stack"
)

func Evaluate(tokens []Token) (operation.Operation, error) {
	s := &stack.Stack[operation.Operation]{}

	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]
		switch t := tok.(type) {
		case Number:
			s.Push(operation.NewNumber(float64(t)))
		case Binary:
			if s.Len() < 2 {
				return nil, fmt.Errorf("no value found for binary operator")
			}

			l := s.Pop()
			r := s.Pop()

			op, err := operation.NewBinary(rune(t), r, l)
			if err != nil {
				return nil, err
			}

			s.Push(op)
		case Function:
			if s.Len() < 1 {
				return nil, fmt.Errorf("no value found for function")
			}

			x := s.Pop()

			f, err := operation.NewFunction(string(t), x)

			if err != nil {
				return nil, err
			}

			s.Push(f)
		}
	}

	if s.IsEmpty() {
		return nil, fmt.Errorf("invalid operation")
	}

	return s.Pop(), nil
}

func TokensToPostfix(tokens []Token) ([]Token, error) {
	s := &stack.Stack[Token]{}
	val := make([]Token, 0)

	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]
		switch tok.(type) {
		case Number:
			val = append(val, tok)
		case Function:
			if len(tokens) <= i+1 {
				return nil, fmt.Errorf("missing left parenthesis after function name")
			}
			_, ok := tokens[i+1].(LParen)
			if !ok {
				return nil, fmt.Errorf("missing left parenthesis after function name")
			}
			i++
			s.Push(tok)
		case LParen:
			s.Push(tok)
		case RParen:
		loop:
			for !s.IsEmpty() {
				t := s.Pop()
				switch t.(type) {
				case Function:
					val = append(val, t)
					break loop
				case LParen:
					break loop
				default:
					val = append(val, t)
				}
			}
		case Binary:
			for !s.IsEmpty() && tok.Pred() <= s.Top().Pred() && tok.Asso() == 'L' {
				val = append(val, s.Pop())
			}
			s.Push(tok)
		}
	}

	for !s.IsEmpty() {
		val = append(val, s.Pop())
	}

	return val, nil
}
