package main

import (
	"ch2/ll"
	"testing"
	"math/rand"
)

// If the list do not intersect, should I return nil?
func TestIntersection(t *testing.T) {
	type test struct {
		l1 *ll.LinkedList[int]
		l2 *ll.LinkedList[int]
		expected *ll.Node[int]
	}
	tests := []test{}
	for i := 0; i < 10; i++ {
		te := test{
			l1: &ll.LinkedList[int]{},
			l2: &ll.LinkedList[int]{},
			expected: ll.NewNode(rand.Intn(100)),
		}
		for j := 0; j < 5; j++ {
			te.l1.AddTail(ll.NewNode(rand.Intn(100)))
			te.l2.AddTail(ll.NewNode(rand.Intn(100)))
		}
		te.l1.AddNodeIdx(i, te.expected)
		te.l2.AddNodeIdx(i, te.expected)
		tests = append(tests, te)
	}
	tests = append(tests, test{nil, nil, nil})
	tests = append(tests, test{
		l1: ll.CreateLinkedList([]int{1, 2}), 
		l2: ll.CreateLinkedList([]int{1, 2}), 
		expected: nil})

	for _, tt := range tests {
		if Intersection(tt.l1, tt.l2) != tt.expected {
			t.Errorf("error Intersection. got=%v, expected=%v\n",
			Intersection(tt.l1, tt.l2), tt.expected)
		}
	}
}
