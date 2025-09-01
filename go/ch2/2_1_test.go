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
	}
}
