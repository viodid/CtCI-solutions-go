package main

import "testing"

func TestIsUnique(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"asdf", true},
		{"asdfqwer12345", true},
		{"asdfa", false},
		{"", true},
		{" ", true},
	}

	for _, tt := range tests {
		if isUnique(tt.input) != tt.expected {
			t.Fatalf("isUnique(%s) does not return %t. got=%t\n",
				tt.input, tt.expected, isUnique(tt.input))
		}
	}
}

func TestIsUniquev2(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"asdf", true},
		{"asdfqwertyuiop", true},
		{"asdfa", false},
		{"", true},
	}

	for _, tt := range tests {
		if isUniquev2(tt.input) != tt.expected {
			t.Fatalf("isUnique(%s) does not return %t. got=%t\n",
				tt.input, tt.expected, isUniquev2(tt.input))
		}
	}
}
