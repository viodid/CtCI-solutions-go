package main

import "testing"

func TestCheckPermutation(t *testing.T) {
	tests := []struct {
		input    []string
		expected bool
	}{
		{[]string{"asdf", "asdf"}, false},
		{[]string{"asdf", "asfd"}, true},
		{[]string{"asdfqwer12345", "123asqwer4df5"}, true},
		{[]string{"asdfa", "asdf"}, false},
		{[]string{"", ""}, false},
		{[]string{" ", " "}, false},
	}

	for _, tt := range tests {
		if checkPermutation(tt.input[0], tt.input[1]) != tt.expected {
			t.Fatalf("IsUnique(%s, %s) does not return %t. got=%t\n",
				tt.input[0], tt.input[1], tt.expected, checkPermutation(tt.input[0], tt.input[1]))
		}
	}
}
