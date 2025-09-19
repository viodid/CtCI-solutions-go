package main

import (
	"ch2/ll"
	"testing"
)

func TestSumLists(t *testing.T) {
	tests := []struct{
		n1 *ll.LinkedList[int]
		n2 *ll.LinkedList[int]
		expected []int
	}{
		{
			ll.CreateLinkedList([]int{1, 2, 1}),
			ll.CreateLinkedList([]int{3, 5}),
			[]int{4, 7, 1},

		},
		{
			ll.CreateLinkedList([]int{9, 2}),
			ll.CreateLinkedList([]int{3, 5}),
			[]int{2, 8},
		},
		{
			ll.CreateLinkedList([]int{9, 4}),
			ll.CreateLinkedList([]int{3, 5}),
			[]int{2, 0, 1},
		},
		{
			ll.CreateLinkedList([]int{9, 4, 0}),
			ll.CreateLinkedList([]int{3, 5}),
			[]int{2, 0, 1},
		},
		{
			ll.CreateLinkedList([]int{0}),
			ll.CreateLinkedList([]int{1}),
			[]int{1},
		},
		{nil, nil, nil},
	}

	for _, tt := range tests {
		output := sumLists(tt.n1, tt.n2)
		if tt.expected == nil || output == nil {
			if tt.expected == nil && output != nil {
				t.Errorf("wrong expected output. expected=%v, got=%v\n",
				tt.expected, output)
			}
			continue
		}
		idx := 0
		for node := output.Head; node != nil; node = node.Next {
			if node.Content != tt.expected[idx] {
				t.Errorf("wrong expected output. expected=%d, got=%d\n",
				tt.expected[idx], node.Content)
			}
			idx++
		}
		if idx != len(tt.expected) {
			t.Errorf("one or more nodes are missing. total nodes=%d, expected=%d\n",
					idx, len(tt.expected))
		}
	}
}
