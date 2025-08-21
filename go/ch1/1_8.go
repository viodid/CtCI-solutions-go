package main

// traverse all the elements, when a 0 is found, store the state
// then set all saved coordinates axis elements to 0.
// time: O(hw + (p * (h + w))) - space: O(p) p = "zero" points
func zeroMatrix(matrix [][]int) [][]int {
	if matrix == nil || len(matrix) == 0 || len(matrix) * len(matrix[0]) <= 1 {
		return matrix
	}
	zeroPoints := [][]int{}
	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == 0 {
				zeroPoints = append(zeroPoints, []int{y, x})
			}
		}
	}
	// p * (h + w)
	for _, points := range zeroPoints {
		setAxisZero(matrix, points[0], points[1])
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

// TODO: O(1) space - use first row and column to store state
