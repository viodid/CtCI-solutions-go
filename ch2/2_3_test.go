package main

import (
	"ch2/ll"
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {
	inputs := []struct {
		nums []int
		idx  int
	}{
		{[]int{3, 2, 1, 0}, 1},
		{[]int{3, 2, 1, 0}, 2},
		{[]int{3, 2, 1, 0, 1}, 3},
		{[]int{3, 1, 0}, 1},
	}

	type test struct {
		list     *ll.LinkedList[int]
		node     *ll.Node[int]
		expected *ll.LinkedList[int]
	}
	tests := []test{}
	for _, in := range inputs {
		list := ll.CreateLinkedList(in.nums)
		node := list.GetNodeIdx(in.idx)
		newList := ll.CreateLinkedList(in.nums)
		newList.DeleteNodeIdx(in.idx)
		tests = append(tests, test{
			list:     list,
			node:     node,
			expected: newList,
		})
	}

	for _, tt := range tests {
		deleteMiddleNode(tt.node)
		if tt.list == nil {
			if tt.expected != nil {
				t.Errorf("deleteMiddleNode failed. got=%+v expected=%v\n",
					tt.list, tt.expected)
			}
			continue
		}
		expNode := tt.expected.Head
		for node := tt.list.Head; node != nil; node = node.Next {
			if expNode == nil {
				t.Errorf("deleteMiddleNode failed. got=%+v expected=%+v\n",
					node, expNode)
				break
			}
			if node.Content != expNode.Content {
				t.Errorf("deleteMiddleNode failed. got=%+v expected=%d\n",
					node.Content, expNode.Content)
			}
			expNode = expNode.Next
		}
	}
}
