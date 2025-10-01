package ch2

import "github.com/viodid/ctci-solutions-go/ll"

// Hashmap, return first collision, else return nil
// time: O(n) - space: O(n)
func LoopDetection(list *ll.LinkedList[int]) *ll.Node[int] {
	if list == nil {
		return nil
	}
	hashNodes := map[*ll.Node[int]]bool{}
	for node := list.Head; node != nil; node = node.Next {
		_, ok := hashNodes[node]
		if ok {
			return node
		}
		hashNodes[node] = true
	}

	return nil
}
