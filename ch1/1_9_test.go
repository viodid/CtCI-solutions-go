package ch1

import "testing"

func TestStringRotation(t *testing.T) {
	tests := []struct {
		input    []string
		expected bool
	}{
		{
			input:    []string{"", ""},
			expected: true,
		},
		{
			input:    []string{"a", "a"},
			expected: true,
		},
		{
			input:    []string{"a", "ab"},
			expected: false,
		},
		{
			input:    []string{"erbottlewat", "waterbottle"},
			expected: true,
		},
		{
			input:    []string{"erottlewat", "waterbottle"},
			expected: false,
		},
	}

	for _, tt := range tests {
		output := stringRotation(tt.input[0], tt.input[1])
		if output != tt.expected {
			t.Errorf("stringRotation(\"%s\", \"%s\") is not %t. got=%t\n",
				tt.input[0], tt.input[1], output, tt.expected)
		}
	}
}
