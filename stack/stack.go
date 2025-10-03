package stack

import "github.com/viodid/ctci-solutions-go/ll"

type Stack[T any] struct {
	data *ll.LinkedList[T]
}

func (s *Stack[T])Pop() T {
	return s.data.RemoveHead().Content
}

func (s *Stack[T])Push(item T) {
	s.data.AddFront(ll.NewNode(t))
}

func (s *Stack[T])Peek() T {
	return s.data.Head.Content
}

func (s *Stack[T])IsEmpty() bool {
	return s.data.Head == nil
}
