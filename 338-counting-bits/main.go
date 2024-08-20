package main

func countBits(n int) []int {
	count := make([]int, n+1)

	for i := 0; i <= n; i++ {
		count[i] = count[i>>1] + (i & 1)
	}

	return count
}
