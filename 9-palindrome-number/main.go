package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	digits := []int{}

	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	for i := 0; i < len(digits); i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}

	return true
}
