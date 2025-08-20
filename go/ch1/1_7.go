package main


// given an x * x size matrix, to rotate it 90 degrees
// the y and x axis should be swapped so the smallest
// rows in x becomes the smallest columns in y, and the biggest
// columns in y becomes the samallest rows in x

// time: O(n) - space: O(n) n = each pixel or input number
func rotateMatrix(image [][]uint8) [][]uint8 {
	if image == nil || len(image) <= 1  {
		return image
	}

	size := len(image)
	matrix := make([][]uint8, size)

	for i := 0; i < size; i++ {
		matrix[i] = make([]uint8, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = image[size - 1 - j][i]
		}
	}

	return matrix
}

// optimization 1
// rotate the matrix in place. big O(1) space
// func rotateMatrixv2(image [][]uint8) [][]uint8 {
// 	if image == nil || len(image) <= 1  {
// 		return image
// 	}
// 	for i := 0; i < size; i++ {
// 		for j := 0; j < math.Ceil(size / 2); j++ {
// 			tmp := image[i][j]
// 			image[i][j] = image[size - 1 - j][i]
// 			image
// 		}
// 	}
// 	return image
// }

// optimization 2
// allocate the matrix as a contiguous 2D array (slighty better big O time)
// time: O(n) - space: O(n)
func rotateMatrixv3(image [][]uint8) [][]uint8 {
	if image == nil || len(image) <= 1  {
		return image
	}

	size := len(image)
	matrix := make([][]uint8, size)
	rows := make([]uint8, size * size)

	for i := 0; i < size; i++ {
		row_start := i * size
		matrix[i] = rows[row_start:row_start + size]
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = image[size - 1 - j][i]
		}
	}

	return matrix
}
