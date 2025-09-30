package main

// time: O(c) c = character set - spcce: O(c)
func isUnique(str string) bool {

	var hashmap = make(map[rune]bool)

	if len(str) >= 128 { // assuming ASCII
		return false
	}

	for _, ch := range str {
		_, ok := hashmap[ch]
		if !ok {
			hashmap[ch] = true
		} else {
			return false
		}
	}
	return true
}

// Asumming only aphabetical ASCII characters
// time: O(c) c = character set - O(1)
func isUniquev2(str string) bool {
	if len(str) >= 26 {
		return false
	}

	var check int
	for _, l := range str {
		val := int(l) - int('a')
		if check&(1<<val) > 0 {
			return false
		}
		check |= (1 << val)
	}
	return true
}
