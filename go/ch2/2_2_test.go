package main

import (
	"ch2/ll"
	"testing"
)

func TestReturnKthToLast(t *testing.T) {
	tests := []struct {
		list *ll.LinkedList[int]
		idx int
		expected *ll.Node[int]
	}{
		{ 
			ll.CreateLinkedList([]int{3, 2, 1, 0}),
			1,
			ll.NewNode(0),
		},
		{ 
			ll.CreateLinkedList([]int{3, 2, 1, 0}),
			2,
			ll.NewNode(1),
		},
		{ 
			ll.CreateLinkedList([]int{3, 2, 1, 0}),
			3,
			ll.NewNode(2),
		},
		{ 
			ll.CreateLinkedList([]int{3, 2, 1, 0}),
			4,
			ll.NewNode(3),
		},
		{ 
			ll.CreateLinkedList([]int{3, 2, 1, 0}),
			0,
			nil,
		},
		{ 
			ll.CreateLinkedList([]int{3, 2, 1, 0}),
			5,
			nil,
		},
		{ 
			nil,
			5,
			nil,
		},
		{ 
			ll.CreateLinkedList([]int{3}),
			1,
			ll.NewNode(3),
		},
	}

	for _, tt := range tests {
		if tt.expected == nil {
			continue
		}
		node := returnKthToLast(tt.list, tt.idx)
		if tt.expected == nil && node != nil {
			t.Errorf("returnKthToLast(nil, %d) does not return expected=%+v\n", tt.idx, tt.expected)
			continue
		}
		if node.Content != tt.expected.Content {
			t.Errorf("returnKthToLast failed. got=%d expected=%d idx=%d\n",
				node.Content, tt.expected.Content, tt.idx)
			continue
		}
	}
}
