package main

import "math"

// this problem can be solved in a couple of different ways
// the first way implies a hashmap and the second an int vector to
// store the state of the char -> repetitions
// as long as both hashmaps are equal or different by one repetition
// the condition is true
// the difference between first and seccond approaches is the big O space time
// and the int bit vector needs to assume ASCII characters only
func oneWay(s1, s2 string) bool {
	s1HashMap := createHashMap(s1)
	s2HashMap := createHashMap(s2)

	if len(s1HashMap) > len(s2HashMap) {
		bigger = s1HashMap
		smaller = s2HashMap
	} else {
		bigger = s2HashMap
		smaller = s1HashMap
	}

	difference := false

	for k, v := range bigger {
		if _, ok := smaller; !ok {
			if v > 1 {
				return false
			} else if v == 1 {
				if difference {
					return false
				}
				difference = true
			}
		}
		valueDiff := math.Abs(v - smaller[k])
		if  valueDiff > 1 {
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
