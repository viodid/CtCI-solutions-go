package main

import (
	"strconv"
	"strings"
)
// the optimal solution for this problem would be to
// construct the final string while traversing the uncompressed string
// when the string is traversed, compare the lengths and return the original
// if it is less than or equal to the compressed string

// for golang, as strings are inmutable, would be better to create an
// slice of strings and append to it so the big O time is not
// increased as every string mutation is a full copy

// time: O(n) - space: O(c) c = character compression repetition
func stringCompression(input string) string {
	compressedStr := []string{}

	input = input + "0"
	lastChar := byte(input[0])
	counter := 1
	for _, ch := range input[1:] {
		if lastChar != byte(ch) {
			compressedStr = append(compressedStr, string(lastChar) + strconv.Itoa(counter))
			counter = 1
		} else {
			counter++
		}
		lastChar = byte(ch)
	}

	input = input[:len(input)-1]

	if len(compressedStr) * 2 < len(input) {
		return strings.Join(compressedStr, "")
	}
	return input
}
