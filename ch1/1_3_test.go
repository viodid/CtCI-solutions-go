package main

import "testing"

func TestURLify(t *testing.T) {
	tests := []struct {
		inputStr    string
		length      int
		expectedStr string
	}{
		{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
		{"Mr John  Smith      ", 14, "Mr%20John%20%20Smith"},
		{" Mr John Smith      ", 14, "%20Mr%20John%20Smith"},
		{"   ", 1, "%20"},
		{"", 0, ""},
		{"hey    there        ", 12, "hey%20%20%20%20there"},
		{"hey there  ", 9, "hey%20there"},
		{"hey", 3, "hey"},
	}

	for _, tt := range tests {
		prevStr := []byte(tt.inputStr)
		URLify(prevStr, tt.length)
		if string(prevStr) != tt.expectedStr {
			t.Errorf("URLify(%s, %d) hasn't mutate the string as '%s'\n",
				prevStr, tt.length, tt.expectedStr)
		}
	}
}
