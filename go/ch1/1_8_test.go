package main

import "testing"

func TestZeroMatrix(t *testing.T) {
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
				{1},
			},
			expected: [][]uint8{
				{1},
			},
		},
		{
			input: [][]uint8{
				{0, 1},
				{2, 3},
			},
			expected: [][]uint8{
				{0, 0},
				{2, 0},
			},
		},
		{
			input: [][]uint8{
				{4, 1},
				{2, 3},
			},
			expected: [][]uint8{
				{4, 1},
				{2, 3},
			},
		},
		{
			input: [][]uint8{
				{10, 0, 69},
				{14, 13, 0},
				{14, 30, 11},
			},
			expected: [][]uint8{
				{0, 0, 0},
				{0, 0, 0},
				{14, 0, 0},
			},
		},
		{
			input: [][]uint8{
				{71, 72, 73, 74},
				{10, 0, 69, 51},
				{15, 13, 68, 52},
				{14, 30, 11, 0},
			},
			expected: [][]uint8{
				{71, 0, 73, 0},
				{0, 0, 0, 0},
				{15, 0, 68, 0},
				{0, 0, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		size := len(tt.input)
		output := zeroMatrix(tt.input)
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if output[i][j] != tt.expected[i][j] {
					t.Errorf("zeroMatrix[%d][%d] is wrong. got=%d. expected=%d\n",
					i, j, output[i][j], tt.expected[i][j])
				}
			}
		}
	}
}
