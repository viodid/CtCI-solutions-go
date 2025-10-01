package ch1

import "testing"

func TestPalindromePermutation(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Tact Coa", true},
		{"aa", true},
		{"a", true},
		{"", false},
		{"David Iva", true},    // ivaddavi
		{"David Viva", true},   // ivadvdavi
		{"David Vivad", false}, // ivaddvdavi
		{"Tact Coa 1 4", true},
		{" 1 4", false},
	}
	for _, tt := range tests {
		if palindromePermutation(tt.input) != tt.expected {
			t.Errorf("palindromePermutation(\"%s\") is not %t. got=%t\n",
				tt.input, tt.expected, palindromePermutation(tt.input))
		}
	}
}

func TestPalindromePermutationv2(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Tact Coa", true},
		{"aa", true},
		{"a", true},
		{"", false},
		{"David Iva", true},    // ivaddavi
		{"David Viva", true},   // ivadvdavi
		{"David Vivad", false}, // ivaddvdavi
		{"Tact Coa 1 4", true},
		{" 1 4", false},
	}
	for _, tt := range tests {
		if palindromePermutationv2(tt.input) != tt.expected {
			t.Errorf("palindromePermutationv2(\"%s\") is not %t. got=%t\n",
				tt.input, tt.expected, palindromePermutationv2(tt.input))
		}
	}
}
