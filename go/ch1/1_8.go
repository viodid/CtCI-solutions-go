package main

// traverse all the elements, when a 0 is found,
// set its y and x elements to 0.
// time: O(n^2) - space: O(1)
func zeroMatrix(matrix [][]int) [][]int {
	if matrix == nil || len(matrix) == 0 || len(matrix) * len(matrix[0]) <= 1 {
		return matrix
	}
	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == 0 {
				setAxisZero(matrix, y, x)
			}
		}
	}
	return matrix
}

func setAxisZero(matrix [][]int, y, x int) {
	height := len(matrix)
	width := len(matrix[0])
	for i := 0; i < width; i++ {
		if matrix[y][i] != 0 {
			matrix[y][i] = 0
		}
	}
	for i := 0; i < height; i++ {
		if matrix[i][x] != 0 {
			matrix[i][x] = 0
		}
	}
}
