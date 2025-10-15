package ch3

import "github.com/viodid/ctci-solutions-go/stack"

type SetOfStacks struct {
	stacks        []stack.Stack[int]
	stackCapacity int
}

func NewSetOfStacks() *SetOfStacks {
	return &SetOfStacks{
		stacks:        []stack.Stack[int]{*stack.NewStack[int]()},
		stackCapacity: 4,
	}
}

func (ss *SetOfStacks) Push(item int) {
	stackIdx := len(ss.stacks) - 1
	currentStack := ss.stacks[stackIdx]
	if currentStack.Length() == 4 {
		ss.stacks = append(ss.stacks, *stack.NewStack[int]())
		currentStack = ss.stacks[stackIdx+1]
	}
	currentStack.Push(item)
}

func (ss *SetOfStacks) Pop() (int, error) {
	currentStack := ss.stacks[len(ss.stacks)-1]
	out, err := currentStack.Pop()
	if currentStack.Length() == 0 && len(ss.stacks) > 1 {
		ss.stacks = ss.stacks[:len(ss.stacks)-1]
	}
	return out, err
}
