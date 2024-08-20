package main

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	curr, prev := 1, 1

	for range n - 1 {
		curr = prev + curr
		prev = curr - prev
	}

	return curr
}
