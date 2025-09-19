package main

import "ch2/ll"

// Iterate through both LL and sum up its digits plus the carry
// create a new node for the addition but module 10
// if the result is bigger than 9, set the carry
// if the lists have different lenghts, append the remaining elemnts
// of the longer list to the output
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
