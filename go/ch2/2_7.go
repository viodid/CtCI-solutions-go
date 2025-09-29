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
