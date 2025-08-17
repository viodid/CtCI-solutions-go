package main

import (
	"strings"
	"unicode"
)

// assumptions: an empty string or non alphabetical is not a palindrome permutation
// a single characeter is a valid palindrome permutation
// unicode characters are vaild characters

// traverse the string and create a hashmap as: character -> occurrences
// traverse the hashmap and check whether all the keys has an even number
// only one key can have an odd number (in case of an odd length permutation)

// time: O(n) - space: O(c) c = character set
func palindromePermutation(input string) bool {
	if len(input) == 0 {
		return false
	}
	charHashMap := make(map[rune]int)
	for _, ch := range strings.ToLower(input) {
		if !unicode.IsLetter(ch) {
			continue
		}
		_, ok := charHashMap[ch]
		if !ok {
			charHashMap[ch] = 1
		} else {
			charHashMap[ch] += 1
		}
	}

	if len(charHashMap) == 0 {
		return false
	}

	checkOdd := false
	for _, v := range charHashMap {
		if v % 2 == 1 {
			if checkOdd == true {
				return false
			}
			checkOdd = true
		}
	}
	return true
}

// for this version, big O space is optimized with an int_32 vector (only 26 possible ASCII alphabetical characters)
// and the character counter can be optimized using a boolean,
// since odd or even is the only property that needs to be stored

// time: O(n) - space: O(1)
func palindromePermutationv2(input string) bool {
	if len(input) == 0 {
		return false
	}

	var intVec int32
	flagIsPermutation := false

	for _, ch := range strings.ToLower(input) {
		byteChar := byte(ch)
		if byteChar < 'a' || byteChar > 'z' {
			continue
		}
		flagIsPermutation = true
		intVec ^= (1 << (byteChar - 'a'))
	}

	// this flags helps to check the edge case in which the string is an even
	// palindrome and hence the int vector is all zeros vs an int vector with all
	// zeros bc no alphabetical character is found
	if !flagIsPermutation {
		return false
	}

	return intVec & (intVec - 1) == 0
}
