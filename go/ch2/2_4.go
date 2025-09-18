package main

import "ch2/ll"

// the easiest way to solve this problem is to
// use two additional data structures to store the
// linked list before and after the pivot
// time: O(n) - space(n)
func partition(list *ll.LinkedList[int], pivot int) {
	if list == nil || list.Head == list.Tail {
		return
	}
	beforePivot := &ll.LinkedList[int]{}
	afterPivot := &ll.LinkedList[int]{}
	for node := list.Head; node != nil; node = node.Next {
		if node.Content < pivot {
			if beforePivot == nil {
				firstNode(node.Content, beforePivot)
				continue
			}
			beforePivot.AddTail(node)
		} else {
			if afterPivot == nil {
				firstNode(node.Content, afterPivot)
				continue
			}
			afterPivot.AddTail(node)
		}
	}
	if beforePivot != nil {
		beforePivot.Tail.Next = afterPivot.Head
	} else {
		beforePivot = afterPivot
	}
	sortedNode := beforePivot.Head
	for node := list.Head; node != nil; node = node.Next {
		node.Content = sortedNode.Content
		sortedNode = sortedNode.Next
	}
}

func firstNode(content int, list *ll.LinkedList[int]) {
	node := &ll.Node[int]{Content: content}
	list.Head = node
	list.Tail = node
}


// the pseudo-sorting could be done in-place leaving constant space
// time: O(n) - space(1)
func partitionv2(list *ll.LinkedList[int]) {
}
