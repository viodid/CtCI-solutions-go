package main

import "ch2/ll"

// time: O(1) - space: O(1)
func deleteMiddleNode(node *ll.Node[int]) {
	if node == nil || node.Next == nil {
		return
	}
	node.Content = node.Next.Content
	node.Next = node.Next.Next
}
