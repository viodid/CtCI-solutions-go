package main

// given an x * x size matrix, to rotate it 90 degrees
// the y and x axis should be swapped so the smallest
// rows in x becomes the smallest columns in y, and the biggest
// columns in y becomes the samallest rows in x

// time: O(n) - space: O(n) n = each pixel or input number
func rotateMatrix(image [][]uint8) [][]uint8 {
	if image == nil || len(image) <= 1 {
		return image
	}

	size := len(image)
	matrix := make([][]uint8, size)

	for i := 0; i < size; i++ {
		matrix[i] = make([]uint8, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = image[size-1-j][i]
		}
	}

	return matrix
}

// optimization 1
// for this optimization, we are going to replace every index in place
// that would make the big O space constant
// rotate every layer (or side) starting from the outtermost layer and
// woring our way inwards. Left -> top -> right -> bottom

// time: O(n^2) n = side length - space: O(1)
func rotateMatrixv2(image [][]uint8) [][]uint8 {
	if image == nil || len(image) <= 1 {
		return image
	}

	size := len(image)
	for layer := 0; layer < size/2; layer++ {
		first := layer
		last := size - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			tmp := image[first][i]                               // top left corner's layer
			image[first][i] = image[last-offset][first]          // top[i] = left[i]
			image[last-offset][first] = image[last][last-offset] // left[i] = bot[i]
			image[last][last-offset] = image[i][last]            // bot[i] = right[i]
			image[i][last] = tmp                                 // right[i] = top[i](tmp)
		}
	}

	return image
}

// optimization 2
// allocate the matrix as a contiguous 2D array (slighty better big O time)
// time: O(n) - space: O(n)
func rotateMatrixv3(image [][]uint8) [][]uint8 {
	if image == nil || len(image) <= 1 {
		return image
	}

	size := len(image)
	matrix := make([][]uint8, size)
	rows := make([]uint8, size*size)

	for i := 0; i < size; i++ {
		row_start := i * size
		matrix[i] = rows[row_start : row_start+size]
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = image[size-1-j][i]
		}
	}

	return matrix
}
