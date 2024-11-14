package parser

import (
	"fmt"

	"github.com/imaah/go-calculator/pkg/operation"
	"github.com/imaah/go-calculator/pkg/stack"
)

func Evaluate(tokens []Token) (operation.Operation, error) {
	s := &stack.Stack[operation.Operation]{}
	commaNumber := 0

	fmt.Println(tokens)
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
		case Unary:
			if s.Len() < 1 {
				return nil, fmt.Errorf("no value found for binary operator")
			}

			r := s.Pop()

			op, err := operation.NewUnary(rune(t), r)
			if err != nil {
				return nil, err
			}

			s.Push(op)
		case Constant:
			f, err := operation.NewConstant(string(t))

			if err != nil {
				return nil, err
			}

			s.Push(f)
		case Function:
			if s.Len() < commaNumber+1 {
				return nil, fmt.Errorf("not enough value found for function")
			}

			values := make([]operation.Operation, commaNumber+1)

			for i := range commaNumber + 1 {
				values[commaNumber-i] = s.Pop()
			}

			f, err := operation.NewFunction2(string(t), values...)
			commaNumber = 0
			if err != nil {
				return nil, err
			}

			s.Push(f)
		case Comma:
			commaNumber++
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
		case Name:
			if len(tokens) <= i+1 {
				val = append(val, Constant(tok.(Name)))
				break
			}
			_, ok := tokens[i+1].(LParen)
			if !ok {
				val = append(val, Constant(tok.(Name)))
				break
			}
			i++
			s.Push(Function(tok.(Name)))
		case LParen:
			s.Push(tok)
		case Comma:
		loopc:
			for !s.IsEmpty() {
				t := s.Top()
				switch t.(type) {
				case Function, Comma:
					break loopc
				default:
					val = append(val, s.Pop())
				}
			}
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
		case Unary:
			s.Push(tok)
		}
	}

	for !s.IsEmpty() {
		val = append(val, s.Pop())
	}

	return val, nil
}
