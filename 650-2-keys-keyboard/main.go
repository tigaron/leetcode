package main

import "math"

func minSteps(n int) int {
	if n <= 1 {
		return 0
	}

	if n%2 == 0 {
		return minSteps(n/2) + 2
	}

	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return minSteps(n/i) + i
		}
	}

	return n
}
