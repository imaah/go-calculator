package stack

import "fmt"

type Stack[T any] struct {
	array []T
}

func (s *Stack[T]) Push(value T) {
	s.array = append(s.array, value)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	v := s.array[len(s.array)-1]
	if len(s.array) == 1 {
		s.array = []T{}
	} else {
		s.array = s.array[:len(s.array)-1]
	}
	return v
}

func (s Stack[T]) Top() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.array[len(s.array)-1]
}

func (s Stack[T]) IsEmpty() bool {
	return len(s.array) == 0
}

func (s Stack[T]) String() string {
	return fmt.Sprint(s.array)
}

func (s Stack[T]) Len() int {
	return len(s.array)
}
