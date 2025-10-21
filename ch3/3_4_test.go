package ch3

import (
	"github.com/viodid/ctci-solutions-go/queue"
	"testing"
)

func TestQueueViaStacks(t *testing.T) {
	queue := queue.NewQueue[int]()
	queueStack := NewQueueViaStacks[int]()

	populateQueues(queue, queueStack)

	for i := 0; i < 10; i++ {
		valQ, _ := queue.Peek()
		valS, err := queueStack.Peek()
		if err != nil {
			t.Fatal(err)
		}
		if valQ != valS {
			t.Errorf("error queueStack Peek. got=%d, expected=%d",
				valS, valQ)
		}
		si, _ := queue.Remove()
		pi, err := queueStack.Remove()
		if err != nil {
			t.Fatal(err)
		}
		if pi != si {
			t.Errorf("error queue Remove. expected=%d, got=%d",
				si, pi)
		}
	}
}

func populateQueues(queue *queue.Queue[int], queueStack *QueueViaStacks[int]) {
	for i := range 10 {
		queue.Add(i)
		queueStack.Add(i)
	}
}
