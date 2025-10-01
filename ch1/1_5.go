package ch1

// this problem can be solved with one hashmap for each string
// as long as both hashmaps are equal or different by one repetition
// the condition is true

// time: O(n) - space: O(c) c = character set
func oneWay(s1, s2 string) bool {
	s1HashMap := createHashMap(s1)
	s2HashMap := createHashMap(s2)

	var bigger map[rune]int
	var smaller map[rune]int
	if len(s1HashMap) > len(s2HashMap) {
		bigger = s1HashMap
		smaller = s2HashMap
	} else {
		bigger = s2HashMap
		smaller = s1HashMap
	}

	difference := false

	for k, v := range bigger {
		if _, ok := smaller[k]; !ok {
			if v > 1 {
				return false
			} else if difference {
				return false
			}
			difference = true
			continue
		}
		valueDiff := abs(v, smaller[k])
		if valueDiff > 1 {
			return false
		} else if valueDiff == 1 {
			if difference {
				return false
			}
			difference = true
		}
	}
	return true
}

func createHashMap(input string) map[rune]int {
	hashMap := map[rune]int{}
	for _, ch := range input {
		_, ok := hashMap[ch]
		if !ok {
			hashMap[ch] = 1
		} else {
			hashMap[ch] += 1
		}
	}
	return hashMap
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
