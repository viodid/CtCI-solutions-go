package ch2

import (
	"github.com/viodid/ctci-solutions-go/ll"
	"math/rand"
	"testing"
)

func TestLoopDetection(t *testing.T) {
	type test struct {
		list     *ll.LinkedList[int]
		expected *ll.Node[int]
	}
	tests := []test{}
	for i := 0; i < 10; i++ {
		te := test{
			list:     &ll.LinkedList[int]{},
			expected: ll.NewNode(rand.Intn(100)),
		}
		for j := 0; j < 8; j++ {
			te.list.AddTail(ll.NewNode(rand.Intn(100)))
		}
		te.list.AddNodeIdx(randRange(4, 8), te.expected)
		te.expected.Next = te.list.GetNodeIdx(randRange(0, 4))
		te.expected = te.expected.Next
		tests = append(tests, te)
	}
	tests = append(tests, test{nil, nil})
	tests = append(tests, test{
		list:     ll.CreateLinkedList([]int{1, 2, 3, 4}),
		expected: nil})

	for _, tt := range tests {
		node := LoopDetection(tt.list)
		if node != tt.expected {
			t.Errorf("error LoopDetection. got=%v, expected=%v\n",
				node, tt.expected)
		}
	}
}
