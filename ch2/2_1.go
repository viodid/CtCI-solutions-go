package ch2

import "github.com/viodid/ctci-solutions-go/ll"

// iterate over the LL and store the repetition state.
// If an int is found more than once, remove the node
// time: O(n) - space: O(n)
func removeDups(ll *ll.LinkedList[int]) *ll.LinkedList[int] {
	if ll == nil || ll.Head == ll.Tail {
		return ll
	}
	prev := ll.Head
	curr := ll.Head.Next
	repetitionState := map[int]bool{prev.Content: true}
	for curr != nil {
		_, ok := repetitionState[curr.Content]
		if ok {
			prev.Next = curr.Next
			curr = curr.Next
			continue
		} else {
			repetitionState[curr.Content] = true
		}
		prev = prev.Next
		curr = curr.Next
	}
	return ll
}

// For this version the constraint is that no buffer is permitted to store state.
// So n^2 comparatinos are performed to remove duplicates
// time: O(n^2) - space: O(1)
func removeDupsv2(ll *ll.LinkedList[int]) *ll.LinkedList[int] {
	if ll == nil || ll.Head.Next == nil {
		return ll
	}
	curr := ll.Head
	for curr != nil && curr.Next != nil {
		prevComp := curr
		comp := curr.Next
		for comp != nil {
			if curr.Content == comp.Content {
				prevComp.Next = comp.Next
				comp = comp.Next
				continue
			}
			prevComp = prevComp.Next
			comp = comp.Next
		}
		curr = curr.Next
	}
	return ll
}
