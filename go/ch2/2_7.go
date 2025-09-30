package main

import "ch2/ll"

// Iterate over both of the lists in n^2 and return the first node
// that exists in both LL
// time: O(n^2) - space: O(1)
func Intersection(l1, l2 *ll.LinkedList[int]) *ll.Node[int] {
	if l1 == nil || l2  == nil {
		return nil
	}
	for node1 := l1.Head; node1 != nil; node1 = node1.Next {
		for node2 := l2.Head; node2 != nil; node2 = node2.Next {
			if node1 == node2 {
				return node1
			}
		}
	}
	return nil
}


// Somethings to bear in mind:
// - If two LL intersect, the last node has to be the same
// - we can chope the longest LL until both of them have the same lenght
// - iterate one step at a time on both LL when the have the same lenght
// time: O(n) - space: O(1)
func Intersectionv2(l1, l2 *ll.LinkedList[int]) *ll.Node[int] {
	if l1 == nil || l2  == nil {
		return nil
	} else if getTail(l1) != getTail(l2) {
		return nil
	}

	// LL cursors
	c1 := l1.Head
	c2 := l2.Head

	// Advance the pointer of the longer LL
	diff := l1.Length() - l2.Length()
	if diff > 0 {
		for i := 0; i < diff; i++ {
			c1 = c1.Next
		}
	} else if diff < 0 {
		for i := 0; i < -diff; i++ {
			c2 = c2.Next
		}
	}
	
	// Traverse until the intersection is found
	for {
		if c1 == c2 {
			return c1
		}
		c1 = c1.Next
		c2 = c2.Next
	}

	return nil
}

func getTail(l *ll.LinkedList[int]) *ll.Node[int] {
	node := l.Head
	for ; node.Next != nil; node = node.Next {}
	return node
}
