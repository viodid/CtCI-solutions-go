package main

import (
	"ch2/ll"
	"testing"
)

func TestRemoveDups(t *testing.T) {
	tests := []struct {
		input    *ll.LinkedList[int]
		expected *ll.LinkedList[int]
	}{
		{
			ll.CreateLinkedList([]int{5, 2, 3, 5}),
			ll.CreateLinkedList([]int{5, 2, 3}),
		},
		{
			ll.CreateLinkedList([]int{5, 2, 3}),
			ll.CreateLinkedList([]int{5, 2, 3}),
		},
		{nil, nil},
		{
			ll.CreateLinkedList([]int{5, 2, 3, 3, 3, 1, 2, 2, 0, 5}),
			ll.CreateLinkedList([]int{5, 2, 3, 1, 0}),
		},
		{
			ll.CreateLinkedList([]int{5, 2, 3, 3, 3}),
			ll.CreateLinkedList([]int{5, 2, 3}),
		},
		{
			ll.CreateLinkedList([]int{5, 2, 3, 3, 3, 3, 3, 3}),
			ll.CreateLinkedList([]int{5, 2, 3}),
		},
		{
			ll.CreateLinkedList([]int{5, 2, 3, 2, 4}),
			ll.CreateLinkedList([]int{5, 2, 3, 4}),
		},
		{
			ll.CreateLinkedList([]int{1, 2, 3, 4}),
			ll.CreateLinkedList([]int{1, 2, 3, 4}),
		},
		{
			ll.CreateLinkedList([]int{0}),
			ll.CreateLinkedList([]int{0}),
		},
	}

	for _, tt := range tests {
		if tt.expected == nil {
			if removeDups(tt.input) != nil {
				t.Errorf("removeDups(nil) does not return expected=%+v\n", tt.expected)
			}
			continue
		}
		ll := removeDups(tt.input)
		expectedNode := tt.expected.Head
		idx := 0
		for node := ll.Head; node != nil; node = node.Next {
			if expectedNode == nil {
				t.Errorf("removeDups failed. got=%d expected=nil idx=%d\n",
					node.Content, idx)
				break
			}
			if node.Content != expectedNode.Content {
				t.Errorf("removeDups failed. got=%d expected=%d idx=%d\n",
					node.Content, expectedNode.Content, idx)
				break
			}
			expectedNode = expectedNode.Next
			idx++
		}
	}
}

func TestRemoveDupsv2(t *testing.T) {
	tests := []struct {
		input    *ll.LinkedList[int]
		expected *ll.LinkedList[int]
	}{
		{
			ll.CreateLinkedList([]int{5, 2, 3, 5}),
			ll.CreateLinkedList([]int{5, 2, 3}),
		},
		{
			ll.CreateLinkedList([]int{5, 2, 3}),
			ll.CreateLinkedList([]int{5, 2, 3}),
		},
		{nil, nil},
		{
			ll.CreateLinkedList([]int{5, 2, 3, 3, 3, 1, 2, 2, 0, 5}),
			ll.CreateLinkedList([]int{5, 2, 3, 1, 0}),
		},
		{
			ll.CreateLinkedList([]int{5, 2, 3, 3, 3}),
			ll.CreateLinkedList([]int{5, 2, 3}),
		},
		{
			ll.CreateLinkedList([]int{5, 2, 3, 3, 3, 3, 3, 3}),
			ll.CreateLinkedList([]int{5, 2, 3}),
		},
		{
			ll.CreateLinkedList([]int{5, 2, 3, 2, 4}),
			ll.CreateLinkedList([]int{5, 2, 3, 4}),
		},
		{
			ll.CreateLinkedList([]int{1, 2, 3, 4}),
			ll.CreateLinkedList([]int{1, 2, 3, 4}),
		},
		{
			ll.CreateLinkedList([]int{0}),
			ll.CreateLinkedList([]int{0}),
		},
	}

	for _, tt := range tests {
		if tt.expected == nil {
			if removeDupsv2(tt.input) != nil {
				t.Errorf("removeDupsv2(nil) does not return expected=%+v\n", tt.expected)
			}
			continue
		}
		ll := removeDupsv2(tt.input)
		expectedNode := tt.expected.Head
		idx := 0
		for node := ll.Head; node != nil; node = node.Next {
			if expectedNode == nil {
				t.Errorf("removeDupsv2 failed. got=%d expected=nil idx=%d\n",
					node.Content, idx)
				break
			}
			if node.Content != expectedNode.Content {
				t.Errorf("removeDupsv2 failed. got=%d expected=%d idx=%d\n",
					node.Content, expectedNode.Content, idx)
				break
			}
			expectedNode = expectedNode.Next
			idx++
		}
	}
}
