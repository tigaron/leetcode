package main

import "unicode"

func lengthOfLastWord(s string) int {
	length := 0
	inWord := false

	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsSpace(rune(s[i])) {
			if inWord {
				break
			}
		} else {
			inWord = true
			length++
		}
	}

	return length
}
