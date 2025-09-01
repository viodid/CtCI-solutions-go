package main

import "testing"

type Node[T any] struct {
	content T
	next *Node[T any]
}

func NewNode[T any](content T) *Node {
	return &Node{content: content}
}

type LinkedList[T any] struct {
	head *Node[T any]
	tail *Node[T any]
}

func NewLinkedList[T any](node *Node[T]) *LinkedList {
	return LinkedList{head: node, tail: node}
}

func (ll *LinkedList[T]) AddTail(node *Node[T]) {
	tail.next = node
}

func (ll *LinkedList[T]) RemoveTail(node *Node[T]) {
	for node := ll.head; node != nil; node.next {
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

func (ll *LinkedList[T]) RemoveHead(node *Node[T]) {
	ll.head = ll.head.next
}

func createLinkedList[T any](data []T) *Node {
	node := Node{content: data[0]}
	for point := range data[1:] {
		node.next := Node{content: point}
		node = node.next
	}
	return head
}

func TestRemoveDups(t *testing.T) {
	tests := []struct{
		input *LinkedList[int]
		expected *LinkedList[int]
	}{
		{
			createLinkedList([]int{5, 2, 3, 5}),
			createLinkedList([]int{5, 2, 3})
		},
	}
}
