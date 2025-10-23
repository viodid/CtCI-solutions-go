package ch3

import (
	"github.com/viodid/ctci-solutions-go/stack"
	"math"
)

func SortStack(s *stack.Stack[int]) error {
	if s.IsEmpty() || s.Length() == 1 {
		return nil
	}
	buffer := stack.NewStack[int]()

	// initialize buffer stack
	val, err := s.Pop()
	if err != nil {
		return err
	}
	buffer.Push(val)

	if err := sortWithStack(s, buffer); err != nil {
		return err
	}

	return nil
}

func sortWithStack(stackA, stackB *stack.Stack[int]) error {
	for !stackA.IsEmpty() {
		valA, err := stackA.Pop()
		if err != nil {
			return err
		}
		valB, err := stackB.Peek()
		if err != nil {
			return err
		}
		if valA < valB {
			if err := pushToStackUntilVal(stackB, stackA, valA); err != nil {
				return err
			}
		}
		stackB.Push(valA)
	}

	if err := pushStack(stackB, stackA); err != nil {
		return err
	}

	return nil
}

func pushToStackUntilVal(src, dst *stack.Stack[int], val int) error {
	currVal := math.MaxInt
	for !src.IsEmpty() && currVal > val {
		currVal, err := src.Pop()
		if err != nil {
			return err
		}
		dst.Push(currVal)
	}
	return nil
}

func pushStack(src, dst *stack.Stack[int]) error {
	for !src.IsEmpty() {
		val, err := src.Pop()
		if err != nil {
			return err
		}
		dst.Push(val)
	}
	return nil
}
