package main

func IsUnique(str string) bool {

	var hashmap = make(map[rune]bool)

	if len(str) == 0 {
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
