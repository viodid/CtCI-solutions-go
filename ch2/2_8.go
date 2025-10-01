package main

import "ch2/ll"

// Hashmap, return first collision, else return nil
// time: O(n) - space: O(n)
func LoopDetection(list *ll.LinkedList[int]) *ll.Node[int] {
	if list == nil {
		return nil
	}
	hashNodes := map[*ll.Node[int]]bool{}
	prev := list.Head
	for node := list.Head.Next; node != nil; node = node.Next {
		_, ok := hashNodes[node]
		if ok {
			return prev 
		}
		hashNodes[node] = true
		prev = prev.Next
	}
	
	return nil
}
