package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	stack := NewQueue[int]()

	if stack.IsEmpty() == false {
		t.Error("error stack should be empty")
	}

	for i := range 10 {
		stack.Add(i)
	}

	item, err := stack.Remove()
	if err != nil {
		t.Fatal(err)
	}
	if item != 0 {
		t.Errorf("error queue pop. expected=%d, got=%d",
			0, item)
	}

	item2, err := stack.Peek()
	if err != nil {
		t.Fatal(err)
	}
	if item2 != 1 {
		t.Errorf("error queue peek. expected=%d, got=%d",
			1, item2)
	}
}
