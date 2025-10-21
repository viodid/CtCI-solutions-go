package ch3

import (
	"github.com/viodid/ctci-solutions-go/stack"
)

type QueueViaStacks[T any] struct {
	stackA *stack.Stack[T]
	stackB *stack.Stack[T]
}

func NewQueueViaStacks[T any]() *QueueViaStacks[T] {
	return &QueueViaStacks[T]{
		stackA: stack.NewStack[T](),
		stackB: stack.NewStack[T](),
	}
}

func (qs *QueueViaStacks[T]) Add(item T) {
	qs.rrebaseStack(qs.stackB, qs.stackA)
	qs.stackA.Push(item)
	qs.rrebaseStack(qs.stackA, qs.stackB)
}

func (qs *QueueViaStacks[T]) rrebaseStack(src, dst *stack.Stack[T]) error {
	for !src.IsEmpty() {
		c, err := src.Pop()
		if err != nil {
			return err
		}
		dst.Push(c)
	}
	return nil
}

func (qs *QueueViaStacks[T]) Remove() (T, error) {
	return qs.stackB.Pop()
}

func (qs *QueueViaStacks[T]) Peek() (T, error) {
	return qs.stackB.Peek()
}

func (qs *QueueViaStacks[T]) IsEmpty() bool {
	return qs.stackB.IsEmpty()
}
