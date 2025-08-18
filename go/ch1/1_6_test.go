package main

import "testing"

func TestStringCompression(t *testing.T) {
	tests := []struct{
		input string
		expected string
	}{
		{ "aabcccccaaaxxxxyyx", "a2b1c5a3x4y2x1" },
		{ "aabcccccaaa", "a2b1c5a3" },
		{ "aabcccaaa", "a2b1c3a3" },
		{ "aabccaaa", "aabccaaa" },
		{ "aabbccaa", "aabbccaa" },
		{ "", "" },
		{ "x", "x" },
		{ "zzz", "z3" },
	}

	for _, tt := range tests {
		if stringCompression(tt.input) != tt.expected {
			t.Errorf("stringCompression(\"%s\") doesn't return %s. got=%s\n",
				tt.input, tt.expected, stringCompression(tt.input))
		}
	}
}
