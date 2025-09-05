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
		node := returnKthToLast(tt.list, tt.idx)
		if tt.expected == nil {
			if  node != nil {
				t.Errorf("returnKthToLast(nil, %d) does not return expected=%+v\n", tt.idx, tt.expected)
			}
			continue
		}
		if node.Content != tt.expected.Content {
			t.Errorf("returnKthToLast failed. got=%d expected=%d idx=%d\n",
				node.Content, tt.expected.Content, tt.idx)
			continue
		}
	}
}

func TestReturnKthToLastv2(t *testing.T) {
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
		node := returnKthToLastv2(tt.list, tt.idx)
		if tt.expected == nil  {
			if node != nil {
				t.Errorf("returnKthToLastv2(%+v, %d) does not return expected=%+v got=%+v\n",
					tt.expected, tt.idx, tt.expected, node)
			}
			continue
		}
		if node == nil {
			t.Errorf("returnKthToLastv2 failed. got=%+v expected=%d idx=%d\n",
				node, tt.expected.Content, tt.idx)
			continue
		}
		if node.Content != tt.expected.Content {
			t.Errorf("returnKthToLastv2 failed. got=%d expected=%d idx=%d\n",
				node.Content, tt.expected.Content, tt.idx)
			continue
		}
	}
}
