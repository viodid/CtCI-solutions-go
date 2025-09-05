package main

import "ch2/ll"

// Get the size of the linked list in one iteration
// Return the corresponding index
// time: O(n) - space: O(1)
func returnKthToLast(list *ll.LinkedList[int], k int) *ll.Node[int] {
	if list == nil {
		return nil
	}

	size := 0
	node := list.Head
	for node != nil {
		size++
		node = node.Next
	}
	node = list.Head
	nodeIdx := size - k
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

// Another way of solving this problem is by using two pointers: p1 and p2
// p1 will advance kth nodes
// iterate through the ll with p1 and p2 at the same pace
// when p2 reaches the end, p1 will be at the right spot
// this improves slightly the time complexity but it is still O(n)
// time: O(n) - space: O(1)
func returnKthToLastv2(list *ll.LinkedList[int], k int) *ll.Node[int] {
	if list == nil {
		return nil
	}
	
	p1 := list.Head
	p2 := list.Head
	i := 1
	for p2 != nil {
		if i == k {
			break
		}
		p2 = p2.Next
		i++
	}
	if p2 == nil {
		return nil
	}
	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

