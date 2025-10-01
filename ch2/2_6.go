package ch2

import "github.com/viodid/ctci-solutions-go/ll"

// because a singly LL can only be traversed in
// one direction, pop all the nodes up to the
// middle and check if the returned numbers
// match with the first half of the LL

// time: O(n) - space: O(n / 2)
func palindrome(list *ll.LinkedList[int]) bool {
	if list == nil {
		return false
	}
	toPop := list.Length() / 2
	for i := 0; i < toPop; i++ {
		node := list.RemoveTail()
		if node.Content != list.GetNodeIdx(i).Content {
			return false
		}
	}
	return true
}
