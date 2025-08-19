package main

import "testing"

func TestRotateMatrix(t *testing.T) {
	tests := []struct{
		input [][]uint8
		expected [][]uint8
	}{
		{
			intput: [][]uint8{
				{0}
			},
			expected: [][]uint8{
				{0}
			},
		},
		{
			intput: [][]uint8{
				{0, 1},
				{2, 3},
			},
			expected: [][]uint8{
				{2, 0},
				{3, 1},
			},
		},
		{
			intput: [][]uint8{
				{10, 42, 69},
				{14, 13, 69},
				{14, 30, 11},
			},
			expected: [][]uint8{
				{14, 14, 10},
				{30, 13, 42},
				{11, 69, 69},
			},
		},
	}

	for _, tt := range tests {
		if rotateMatrix(tt.input) != tt.expected {
			t.Errorf("rotateMatrix(%v) hasn't returned %v\n",
				tt.input, tt.expected)
		}
	}
}
