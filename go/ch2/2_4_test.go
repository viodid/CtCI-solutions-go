package main

import (
	"testing"
	"ch2/ll"
)

func TestPartition(t *testing.T) {

	tests := []struct{
		list *ll.LinkedList[int]
		pivot int
	}{
		{ll.CreateLinkedList([]int{3, 5, 8, 5, 10, 2, 1}), 5},
		{ll.CreateLinkedList([]int{3, 5, 8, 5, 10, 2, 1}), 3},
		{ll.CreateLinkedList([]int{3, 5, 8, 5, 10, 2, 1}), 1},
		{ll.CreateLinkedList([]int{3, 5, 8, 5, 10, 2, 1}), 10},
		{ll.CreateLinkedList([]int{1}), 1},
		{ll.CreateLinkedList([]int{1, 2}), 3}, // edge
		{nil, 3}, // edge
	}

	for _, tt := range tests {
		beforePivot := true
		partition(tt.list)
		for node := tt.list.Head; node != nil; node = node.Next {
			if beforePivot && node.Content > tt.pivot {
				t.ErrorF("partition error before pivot. got=%d, pivot=%d\n",
				node.Content, tt.pivot)
				break
			} else if !beforePivot && node.Content <= tt.pivot {
				t.ErrorF("partition error after pivot. got=%d, pivot=%d\n",
				node.Content, tt.pivot)
				break
			}
		}
	}
}
