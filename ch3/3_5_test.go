package ch3

import (
	"github.com/viodid/ctci-solutions-go/stack"
	"math/rand"
	"testing"
)

func TestSortStack(t *testing.T) {
	stack := stack.NewStack[int]()

	for i := 0; i < 10; i++ {
		stack.Push(rand.Int() % 10)
	}

	SortStack(stack)

	prv := -1
	c, err := stack.Pop()
	if err != nil {
		t.Fatal(err)
	}
	for !stack.IsEmpty() {
		if c < prv {
			t.Errorf("stack is not sorted. prev=%d, curr=%d", prv, c)
		}
		prv = c
		c, err = stack.Pop()
		if err != nil {
			t.Fatal(err)
		}
	}
}
