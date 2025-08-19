package main

import "testing"

func TestRotateMatrix(t *testing.T) {
	tests := []struct{
		input [][]uint8
		expected [][]uint8
	}{
		{
			input: [][]uint8{
				{0},
			},
			expected: [][]uint8{
				{0},
			},
		},
		{
			input: [][]uint8{
				{0, 1},
				{2, 3},
			},
			expected: [][]uint8{
				{2, 0},
				{3, 1},
			},
		},
		{
			input: [][]uint8{
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
		size := len(tt.input)
		output := rotateMatrix(tt.input)
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if output[i][j] != tt.expected[i][j] {
					t.Errorf("rotateMatrix[%d][%d] is wrong. got=%d. expected=%d\n",
					i, j, output[i][j], tt.expected[i][j])
				}
			}
		}
	}
}
