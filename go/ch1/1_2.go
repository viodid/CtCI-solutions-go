package main

// assumptions: a permutation of a string should always be different meaning
// the character set is unicode
// when cannot store an occurrence bigger than 2^32
// that the characters cannot be placed in the same order
// guards: check if string is empty, s1 == s2
// create hashmap to store a counter for each character for the s1
// create a second hashmap with the s2
// if both hashmaps have the same ocurrences, true else false

// optimization 1: use only one hashmap and decretment the occurrences
// any given value below 0 is not permutation
// at the end all the occurrences should be 0 to be a permutation

// optimization 2: use an array to store the value counter and the index to 
// store the character (this assumes ASCII or relative)

// time: O(n) - space: O(c) c=character set
func checkPermutation(s1, s2 string) bool {
	if len(s1) == 0 || s1 == s2 || len(s1) != len(s2) {
		return false
	}

	s1Hashmap := make(map[rune]int)
	s2Hashmap := make(map[rune]int)
	for _, l := range s1 {
		_, ok := s1Hashmap[l]
		if !ok {
			s1Hashmap[l] = 1
		} else {
			s1Hashmap[l] += 1
		}
	}
	for _, l := range s2 {
		_, ok := s2Hashmap[l]
		if !ok {
			s2Hashmap[l] = 1
		} else {
			s2Hashmap[l] += 1
		}
	}

	for k, v := range s1Hashmap {
		oc, ok := s1Hashmap[k]
		if !ok || oc != v {
			return false
		}
	}
	return true
}
