package main

import (
	"ch2/ll"
	"testing"
)

func TestPalindrome(t *testing.T) {
	tests := []struct {
		input    *ll.LinkedList[int]
		expected bool
	}{
		{
			ll.CreateLinkedList([]int{1, 2, 1}),
			true,
		},
		{
			ll.CreateLinkedList([]int{1, 2, 1, 2}),
			false,
		},
		{
			ll.CreateLinkedList([]int{1}),
			true,
		},
		{
			ll.CreateLinkedList([]int{1, 2, 3, 2, 5}),
			false,
		},
		{
			ll.CreateLinkedList([]int{1, 2, 2, 1}),
			true,
		},
		{nil, false},
	}

	for idx, tt := range tests {
		if palindrome(tt.input) != tt.expected {
			t.Errorf("wrong output. expected=%t, got=%t, idx=%d\n",
				tt.expected, palindrome(tt.input), idx)
		}
	}
}
