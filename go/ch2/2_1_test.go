package main

import "testing"

type Node[T any] struct {
	content T
	next *Node[T]
}

func NewNode[T any](content T) *Node[T] {
	return &Node[T]{content: content}
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func NewLinkedList[T any](node *Node[T]) *LinkedList[T] {
	return &LinkedList[T]{head: node, tail: node}
}

func (ll *LinkedList[T]) AddTail(node *Node[T]) {
	ll.tail.next = node
	ll.tail = node
}

func (ll *LinkedList[T]) RemoveTail(node *Node[T]) {
	// only one node left
	if ll.head == ll.tail {
		ll.head = nil
		ll.tail= nil
		return
	}
	// only two nodes left
	if ll.head.next == ll.tail {
		ll.head.next = nil
		ll.tail = ll.head
		return
	}
	for node := ll.head; node != nil; node = node.next {
		if node.next.next == nil {
			node.next = nil
			ll.tail = node
			break
		}
	}
}

func (ll *LinkedList[T]) AddFront(node *Node[T]) {
	node.next = ll.head
	ll.head = node
}

func (ll *LinkedList[T]) RemoveHead() {
	if ll.head == nil {
		return
	}
	ll.head = ll.head.next
}

func createLinkedList[T any](data []T) *LinkedList[T] {
	ll := NewLinkedList(NewNode(data[0]))
	for _, point := range data[1:] {
		ll.AddTail(NewNode(point))
	}
	return ll
}

func TestRemoveDups(t *testing.T) {
	tests := []struct{
		input *LinkedList[int]
		expected *LinkedList[int]
	}{
		{
			createLinkedList([]int{5, 2, 3, 5}),
			createLinkedList([]int{5, 2, 3}),
		},
		{
			createLinkedList([]int{5, 2, 3}),
			createLinkedList([]int{5, 2, 3}),
		},
		{ nil, nil },
		{
			createLinkedList([]int{5, 2, 3, 3, 3, 1, 2, 2, 0, 5}),
			createLinkedList([]int{5, 2, 3, 1, 0}),
		},
		{
			createLinkedList([]int{5, 2, 3, 3, 3}),
			createLinkedList([]int{5, 2, 3}),
		},
		{
			createLinkedList([]int{5, 2, 3, 3, 3, 3, 3, 3}),
			createLinkedList([]int{5, 2, 3}),
		},
		{
			createLinkedList([]int{5, 2, 3, 2, 4}),
			createLinkedList([]int{5, 2, 3, 4}),
		},
		{
			createLinkedList([]int{1, 2, 3, 4}),
			createLinkedList([]int{1, 2, 3, 4}),
		},
		{
			createLinkedList([]int{0}),
			createLinkedList([]int{0}),
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
		expectedNode := tt.expected.head
		idx := 0
		for node := ll.head; node != nil; node = node.next {
			if expectedNode == nil {
				t.Errorf("removeDups failed. got=%d expected=nil idx=%d\n",
				node.content, idx)
				break
			}
			if node.content != expectedNode.content {
				t.Errorf("removeDups failed. got=%d expected=%d idx=%d\n",
				node.content, expectedNode.content, idx)
				break
			}
			expectedNode = expectedNode.next
			idx++
		}
	}
}

func TestRemoveDupsv2(t *testing.T) {
	tests := []struct{
		input *LinkedList[int]
		expected *LinkedList[int]
	}{
		{
			createLinkedList([]int{5, 2, 3, 5}),
			createLinkedList([]int{5, 2, 3}),
		},
		{
			createLinkedList([]int{5, 2, 3}),
			createLinkedList([]int{5, 2, 3}),
		},
		{ nil, nil },
		{
			createLinkedList([]int{5, 2, 3, 3, 3, 1, 2, 2, 0, 5}),
			createLinkedList([]int{5, 2, 3, 1, 0}),
		},
		{
			createLinkedList([]int{5, 2, 3, 3, 3}),
			createLinkedList([]int{5, 2, 3}),
		},
		{
			createLinkedList([]int{5, 2, 3, 3, 3, 3, 3, 3}),
			createLinkedList([]int{5, 2, 3}),
		},
		{
			createLinkedList([]int{5, 2, 3, 2, 4}),
			createLinkedList([]int{5, 2, 3, 4}),
		},
		{
			createLinkedList([]int{1, 2, 3, 4}),
			createLinkedList([]int{1, 2, 3, 4}),
		},
		{
			createLinkedList([]int{0}),
			createLinkedList([]int{0}),
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
		expectedNode := tt.expected.head
		idx := 0
		for node := ll.head; node != nil; node = node.next {
			if expectedNode == nil {
				t.Errorf("removeDupsv2 failed. got=%d expected=nil idx=%d\n",
				node.content, idx)
				break
			}
			if node.content != expectedNode.content {
				t.Errorf("removeDupsv2 failed. got=%d expected=%d idx=%d\n",
				node.content, expectedNode.content, idx)
				break
			}
			expectedNode = expectedNode.next
			idx++
		}
	}
}

