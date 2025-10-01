package main

// constraints: time complexity should be at most O(n)

// traverse the input slice backwards with two pointers:
// frist starting from the tLength idx without the right-hand buffer offset
// as the read cursor
// second from the end of the slice as the write cursor

// time: O(n) - space: O(1)
func URLify(inputStr []byte, tLength int) {
	if inputStr == nil || len(inputStr) == 0 {
		return
	}
	writeCur := len(inputStr) - 1

	for readCur := tLength - 1; readCur >= 0; readCur-- {
		if inputStr[readCur] == ' ' {
			inputStr[writeCur] = '0'
			writeCur--
			inputStr[writeCur] = '2'
			writeCur--
			inputStr[writeCur] = '%'
			writeCur--
		} else {
			inputStr[writeCur] = inputStr[readCur]
			writeCur--
		}
	}

}
