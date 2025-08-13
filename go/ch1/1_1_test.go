package main

import "testing"

func TestIsUnique(t *testing.T) {
	tests := []struct{
		input string
		expected bool
	}{
		{ "asdf", true },
		{ "asdfqwer12345", true },
		{ "asdfa", false },
		{ "", false },
		{ " ", true },
	}

	for _, tt := range tests {
		if IsUnique(tt.input) != tt.expected {
			t.Fatalf("IsUnique(%s) does not return %t. got=%t\n",
				tt.input, tt.expected, IsUnique(tt.input))
		}
	}
}
