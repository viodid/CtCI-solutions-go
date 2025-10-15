package stack

import (
	"errors"
	"github.com/viodid/ctci-solutions-go/ll"
)

type Stack[T any] struct {
	data *ll.LinkedList[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: &ll.LinkedList[T]{}}
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("cannot pop on an empty stack")
	}
	return s.data.RemoveHead().Content, nil
}

func (s *Stack[T]) PopAt(idx int) (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("cannot pop on an empty stack")
	}
	if idx >= s.data.Length() || idx < 0 {
		var zero T
		return zero, errors.New("index is out of bounds")
	}
	node := s.data.GetNodeIdx(idx)
	s.data.DeleteNodeIdx(idx)

	return node.Content, nil
}


func (s *Stack[T]) Push(item T) {
	s.data.AddFront(ll.NewNode(item))
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("cannot peek on an empty stack")
	}
	return s.data.Head.Content, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.data == nil || s.data.Head == nil
}

func (s *Stack[T]) Length() int {
	return s.data.Length()
}
