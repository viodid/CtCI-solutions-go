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
