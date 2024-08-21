package main

import "strings"

func lengthOfLastWord(s string) int {
	sub := strings.Fields(s)
	return len(sub[len(sub)-1])
}
