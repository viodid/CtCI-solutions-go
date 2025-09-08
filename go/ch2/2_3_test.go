package main

import (
	"testing"
	"ch2/ll"
)

func TestDeleteMiddleNode(t *testing.T) {

	inputs := []struct{
		nums []int
		idx int
	}{
		{ {3, 2, 1, 0}, 2 },
		{ {3, 2, 1, 0}, 3 },
		{ {3, 2, 1, 0}, 10 },
		{ {3, 1, 0}, 1 },
		{ {3, 1}, 1 },
		{ nil, 1 },
	}

	type test struct {
		list *ll.LinkedList[int]
		node *ll.Node[int]
		expected *ll.LinkedList[int]
	}
	tests := []test{}
	for _, in := range inputs {
		list := ll.CreateLinkedList(in.nums)
		node := list.GetNodeIdx(in.idx)
		expected := list.DeleteNodeIdx(in.idx)
		tests = append(tests, {
			list: list,
			node: node,
			expected: expected,
		})
	}

	for _, tt := range tests {
		// TODO
	}
}
