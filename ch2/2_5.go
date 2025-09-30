package main

import "ch2/ll"

// Iterate through both LL and sum up its digits plus the carry
// create a new node for the addition but module 10
// if the result is bigger than 9, set the carry
// if the lists have different lenghts, append the remaining elemnts
// of the longer list to the output
//
// time: O(a + b) - space: O(max(a, b))
func sumLists(n1, n2 *ll.LinkedList[int]) *ll.LinkedList[int] {
	if n1 == nil || n2 == nil {
		return nil
	} else if n1 == nil {
		return n2
	} else if n2 == nil {
		return n1
	}
	carry := 0
	node1 := n1.Head
	node2 := n2.Head
	output := &ll.LinkedList[int]{}
	for node1 != nil && node2 != nil {
		addition := node1.Content + node2.Content + carry
		if carry == 1 {
			carry = 0
		}
		output.AddTail(ll.NewNode(addition % 10))
		if addition > 9 {
			carry = 1
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	if node1 == nil && node2 == nil {
		if carry == 1 {
			output.AddTail(ll.NewNode(1))
		}
		return output
	}
	var node *ll.Node[int]
	if node1 == nil {
		node = node2
	} else {
		node = node1
	}
	for ;node != nil; node = node.Next {
		output.AddTail(ll.NewNode(node.Content + carry))
		if carry == 1 {
			carry = 0
		}
	}
	return output
}

// follow up: the digits are stored in forward order

// as we need to sum up the digits for the least significant (the last node)
// to the most significant, recursion makes sense
//
// add padding with zeroes if one list is shorter, so both list are aligned
// traverse recursively the linked lists until both are nil.
// as the returned linked list is shared via reference, we can add nodes
// despite the call stack int which it's executing
// if the addition is greater than 9, return the carry
func sumListsv2(n1, n2 *ll.LinkedList[int]) *ll.LinkedList[int] {
	if n1 == nil || n2 == nil {
		return nil
	} else if n1 == nil {
		return n2
	} else if n2 == nil {
		return n1
	}
	result := &ll.LinkedList[int]{}
	equalLenLists(n1, n2)
	if recursiveSumLists(n1.Head, n2.Head, result) == 1 {
		result.AddFront(ll.NewNode(1))
	}
	return result
}

func recursiveSumLists(n1, n2 *ll.Node[int], result *ll.LinkedList[int]) int {
	if n1 == nil && n2 == nil {
		return 0
	}
	carry := recursiveSumLists(n1.Next, n2.Next, result)
	addition := n1.Content + n2.Content + carry
	result.AddFront(ll.NewNode(addition % 10))
	outCarry := 0
	if addition > 9 {
		outCarry = 1
	}
	return outCarry
}

func equalLenLists(n1, n2 *ll.LinkedList[int]) {
	if n1.Length() > n2.Length() {
		paddingZeroList(n2, n1.Length())
	} else if n2.Length() > n1.Length() {
		paddingZeroList(n1, n2.Length())
	}
}

func paddingZeroList(list *ll.LinkedList[int], len int) {
	diff := len - list.Length()
	for i := 0; i < diff; i++ {
		list.AddFront(ll.NewNode(0))
	}
}
