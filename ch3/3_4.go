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

func (qs *QueueViaStacks[T]) Add(item T) error {
	if err := qs.pushStack(qs.stackB, qs.stackA); err != nil {
		return err
	}
	qs.stackA.Push(item)
	return nil
}

func (qs *QueueViaStacks[T]) pushStack(src, dst *stack.Stack[T]) error {
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
	if err := qs.pushStack(qs.stackA, qs.stackB); err != nil {
		var zero T
		return zero, err
	}
	return qs.stackB.Pop()
}

func (qs *QueueViaStacks[T]) Peek() (T, error) {
	if err := qs.pushStack(qs.stackA, qs.stackB); err != nil {
		var zero T
		return zero, err
	}
	return qs.stackB.Peek()
}

func (qs *QueueViaStacks[T]) IsEmpty() (bool, error) {
	if err := qs.pushStack(qs.stackA, qs.stackB); err != nil {
		var zero bool
		return zero, err
	}
	return qs.stackB.IsEmpty(), nil
}
