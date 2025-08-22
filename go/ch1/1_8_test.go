package main

import "testing"

var TESTS = []struct{
	input [][]int
	expected [][]int
}{
	{
		input: nil,
		expected: nil,
	},
	{
		input: [][]int{},
		expected: [][]int{},
	},
	{
		input: [][]int{
			{0},
		},
		expected: [][]int{
			{0},
		},
	},
	{
		input: [][]int{
			{1, 0},
		},
		expected: [][]int{
			{0, 0},
		},
	},
	{
		input: [][]int{
			{0, 1},
			{2, 3},
		},
		expected: [][]int{
			{0, 0},
			{0, 3},
		},
	},
	{
		input: [][]int{
			{4, 1},
			{2, 3},
		},
		expected: [][]int{
			{4, 1},
			{2, 3},
		},
	},
	{
		input: [][]int{
			{10, 0, 69},
			{14, 13, 0},
			{14, 30, 11},
			{19, 15, 18},
		},
		expected: [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{14, 0, 0},
			{19, 0, 0},
		},
	},
	{
		input: [][]int{
			{71, 72, 73, 74},
			{10, 0, 69, 51},
			{15, 13, 68, 52},
			{14, 30, 11, 0},
		},
		expected: [][]int{
			{71, 0, 73, 0},
			{0, 0, 0, 0},
			{15, 0, 68, 0},
			{0, 0, 0, 0},
		},
	},
}

func TestZeroMatrix(t *testing.T) {
	for _, tt := range TESTS {
		if tt.input == nil {
			if zeroMatrix(tt.input) != nil {
				t.Errorf("zeroMatrix(%v) is wrong. got=%v. expected=%v\n",
				tt.input, zeroMatrix(tt.input), tt.expected)
			}
			continue
		}
		if len(tt.input) == 0 {
			if len(zeroMatrix(tt.input)) != 0 {
				t.Errorf("zeroMatrix(%v) is wrong. got=%v. expected=%v\n",
				tt.input, zeroMatrix(tt.input), tt.expected)
			}
			continue
		}
		height := len(tt.input)
		width := len(tt.input[0])
		output := zeroMatrix(tt.input)
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if output[i][j] != tt.expected[i][j] {
					t.Errorf("zeroMatrix[%d][%d] is wrong. got=%d. expected=%d\n",
					i, j, output[i][j], tt.expected[i][j])
				}
			}
		}
	}
}


func TestZeroMatrixv2(t *testing.T) {
	for _, tt := range TESTS {
		if tt.input == nil {
			if zeroMatrix(tt.input) != nil {
				t.Errorf("zeroMatrixv2(%v) is wrong. got=%v. expected=%v\n",
				tt.input, zeroMatrixv2(tt.input), tt.expected)
			}
			continue
		}
		if len(tt.input) == 0 {
			if len(zeroMatrixv2(tt.input)) != 0 {
				t.Errorf("zeroMatrixv2(%v) is wrong. got=%v. expected=%v\n",
				tt.input, zeroMatrixv2(tt.input), tt.expected)
			}
			continue
		}
		height := len(tt.input)
		width := len(tt.input[0])
		output := zeroMatrixv2(tt.input)
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if output[i][j] != tt.expected[i][j] {
					t.Errorf("zeroMatrixv2[%d][%d] is wrong. got=%d. expected=%d\n",
					i, j, output[i][j], tt.expected[i][j])
				}
			}
		}
	}
}
