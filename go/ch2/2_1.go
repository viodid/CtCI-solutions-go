package main

// iterate over the LL and store the repetition state.
// If an int is found more than once, remove the node
// time: O(n) - space: O(n)
func removeDups(ll *LinkedList[int]) *LinkedList[int] {
	if ll == nil || ll.head == ll.tail {
		return ll
	}
	prev := ll.head
	curr := ll.head.next
	repetitionState := map[int]bool{prev.content: true}
	for curr != nil {
		_, ok := repetitionState[curr.content]
		if ok {
			prev.next = curr.next
			curr = curr.next
			continue
		} else {
			repetitionState[curr.content] = true
		}
		prev = prev.next
		curr = curr.next
	}
	return ll
}

// for this version the constraint is that no buffer is permitted to store state
// so when the LL is sorted, is enough to keep just adjacent numbers to delete duplicates
// time: O(nlogn) - space: O(1)
// func removeDupsv2(ll *LinkedList[int]) *LinkedList[int] {
// 	// TODO
// }
