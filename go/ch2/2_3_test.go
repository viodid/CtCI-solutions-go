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
		{ []int{3, 2, 1, 0}, 1 },
		{ []int{3, 2, 1, 0}, 2 },
		{ []int{3, 2, 1, 0}, 10 },
		{ []int{3, 1, 0}, 1 },
		{ []int{3, 1}, 1 },
		// { nil, 1 },
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
		newList := ll.CreateLinkedList(in.nums)
		newList.DeleteNodeIdx(in.idx)
		tests = append(tests, test{
			list: list,
			node: node,
			expected: newList,
		})
	}

	for _, tt := range tests {
		deleteMiddleNode(tt.node)
		expNode := tt.expected.Head
		for node := tt.list.Head; node != nil; node = node.Next {
			if node.Content != expNode.Content {
				t.Errorf("deleteMiddleNode failed. got=%+v expected=%d\n",
					node.Content, expNode.Content)
			}
			expNode = expNode.Next
		}
	}
}
