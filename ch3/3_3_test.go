package ch3

import (
	"github.com/viodid/ctci-solutions-go/stack"
	"testing"
)

func TestStackOfPlates(t *testing.T) {
	stack := stack.NewStack[int]()
	plates := NewSetOfStacks()

	for i := range 10 {
		stack.Push(i)
		plates.Push(i)
	}

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
}
