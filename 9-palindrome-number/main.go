package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	copy := x
	var check int

	for x > 0 {
		digit := x % 10
		check = (check * 10) + digit
		x /= 10
	}

	return copy == check
}
