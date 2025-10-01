package ch1

import "testing"

func TestOneWay(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "bales", false},
		{"pale", "bale", true},
		{"pale", "bake", false},
		{"", "", true},
		{"pale", "pale", true},
		{"a", "", true},
		{"ab", "", false},
		{"hey there", "hi", false},
	}

	for _, tt := range tests {
		if oneWay(tt.s1, tt.s2) != tt.expected {
			t.Errorf("oneWay(\"%s\", \"%s\") is not %t. got=%t",
				tt.s1, tt.s2, tt.expected, oneWay(tt.s1, tt.s2))
		}
	}

}
