package ch1

// traverse all the elements, when a 0 is found, store the state
// then set all saved coordinates axis elements to 0.
// time: O(hw + (p * (h + w))) - space: O(p) p = "zero" points
func zeroMatrix(matrix [][]int) [][]int {
	if matrix == nil || len(matrix) == 0 || len(matrix)*len(matrix[0]) <= 1 {
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

// time: O(hw) - O(1) space - use first row and column to store state
func zeroMatrixv2(matrix [][]int) [][]int {
	if matrix == nil || len(matrix) == 0 || len(matrix)*len(matrix[0]) <= 1 {
		return matrix
	}
	columnHasZero := false
	rowHasZero := false
	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == 0 {
				if y == 0 || x == 0 {
					if !rowHasZero {
						rowHasZero = y == 0
					}
					if !columnHasZero {
						columnHasZero = x == 0
					}
					continue
				}
				matrix[0][x] = 0
				matrix[y][0] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			nullifyRow(matrix, i)
		}
	}
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			nullifyColumn(matrix, i)
		}
	}
	if columnHasZero {
		nullifyColumn(matrix, 0)
	}
	if rowHasZero {
		nullifyRow(matrix, 0)
	}
	return matrix
}

func nullifyRow(matrix [][]int, y int) {
	width := len(matrix[0])
	for i := 0; i < width; i++ {
		if matrix[y][i] != 0 {
			matrix[y][i] = 0
		}
	}
}

func nullifyColumn(matrix [][]int, x int) {
	height := len(matrix)
	for i := 0; i < height; i++ {
		if matrix[i][x] != 0 {
			matrix[i][x] = 0
		}
	}
}
