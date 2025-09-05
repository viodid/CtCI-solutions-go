package main

import "ch2/ll"

// Get the size of the linked list in one iteration
// Return the corresponding index
// time: O(n) - space: O(1)
func returnKthToLast(ll *ll.LinkedList[int], idx int) *ll.Node[int] {
	if ll == nil {
		return nil
	} else if ll.Head.Next == nil && idx == 1 {
		return ll.Head
	}

	size := 0
	node := ll.Head
	for node != nil {
		size++
		node = node.Next
	}
	node = ll.Head
	nodeIdx := size - idx
	cursor := 0
	for node != nil {
		if cursor == nodeIdx {
			return node
		}
		cursor++
		node = node.Next
	}
	return nil
}
