package main

// time: O(a) a = character set - space: O(a)
func IsUnique(str string) bool {

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
