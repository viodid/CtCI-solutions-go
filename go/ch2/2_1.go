package main

// iterate over the LL and store the repetition state.
// If an int is found more than once, remove the node
// time: O(n) - space: O(n)
func removeDups(ll *LinkedList[int]) *LinkedList[int] {
	if ll == nil || ll.head == ll.tail {
		return ll
	}
	n1 := ll.head
	n2 := ll.head.next
	repetitionState := map[int]bool{n1.content: true}
	for n2 != nil {
		_, ok := repetitionState[n2.content]
		if ok {
			n1.next = n2.next
		} else {
			repetitionState[n2.content] = true
		}
		n1 = n1.next
		n2 = n2.next
	}
	return ll
}

// for this version the constraint is that no buffer is permitted to store state
// so when the LL is sorted, is enough to keep just adjacent numbers to delete duplicates
// time: O(nlogn) - space: O(1)
// func removeDupsv2(ll *LinkedList[int]) *LinkedList[int] {
// 	// TODO
// }
