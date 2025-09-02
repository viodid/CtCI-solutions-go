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

// For this version the constraint is that no buffer is permitted to store state.
// So n^2 comparatinos are performed to remove duplicates
// time: O(n^2) - space: O(1)
func removeDupsv2(ll *LinkedList[int]) *LinkedList[int] {
	if ll == nil || ll.head.next == nil {
		return ll
	}
	curr := ll.head
	for curr != nil && curr.next != nil {
		prevComp := curr
		comp := curr.next
		for comp != nil {
			if curr.content == comp.content {
				prevComp.next = comp.next
				comp = comp.next
				continue
			}
			prevComp = prevComp.next
			comp = comp.next
		}
		curr = curr.next
	}
	return ll
}
