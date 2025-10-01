package main

import "strings"

// time: O(n) - space: O(1)
func stringRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return strings.Contains(s1+s1, s2)
}
