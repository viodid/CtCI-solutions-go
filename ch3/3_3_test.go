package ch3

import (
	"github.com/viodid/ctci-solutions-go/stack"
	"testing"
)

func TestStackOfPlates(t *testing.T) {
	stack := stack.NewStack[int]()
	plates := NewSetOfStacks()

	populateStacks(stack, plates)

	for i := 0; i < 10; i++ {
		si, _ := stack.Pop()
		pi, err := plates.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if pi != si {
			t.Errorf("error stack pop. expected=%d, got=%d",
				si, pi)
		}
	}

	populateStacks(stack, plates)

	popPlates(t, plates, 0, 3)
	popPlates(t, plates, 0, 2)
	popPlates(t, plates, 2, 9)
	popPlates(t, plates, 0, 1)
	popPlates(t, plates, 0, 0)
	popPlates(t, plates, 0, 7)
}

func popPlates(t *testing.T, plates *SetOfStacks, idx, expected int) {
	pi, err := plates.PopAt(idx)
	if err != nil {
		t.Fatal(err)
	}
	if pi != expected {
		t.Errorf("error stack PopAt. expected=%d, got=%d", expected, pi)
	}
}

func populateStacks(stack *stack.Stack[int], plates *SetOfStacks) {
	for i := range 10 {
		stack.Push(i)
		plates.Push(i)
	}
}
