package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack[int]()

	if stack.IsEmpty() == false {
		t.Error("error stack should be empty")
	}

	for i := range 10 {
		stack.Push(i)
	}

	item, err := stack.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if item != 9 {
		t.Errorf("error stack pop. expected=%d, got=%d",
			9, item)
	}

	item2, err := stack.Peek()
	if err != nil {
		t.Fatal(err)
	}
	if item2 != 8 {
		t.Errorf("error stack peek. expected=%d, got=%d",
			8, item2)
	}
	item3, err := stack.PopAt(0)
	if err != nil {
		t.Fatal(err)
	}
	if item3 != 8 {
		t.Errorf("error stack peek. expected=%d, got=%d",
			8, item3)
	}
}
