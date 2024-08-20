package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	count := 0
	stepOne := 1
	stepTwo := 2

	for i := 3; i <= n; i++ {
		count = stepOne + stepTwo
		stepOne = stepTwo
		stepTwo = count
	}

	return count
}
