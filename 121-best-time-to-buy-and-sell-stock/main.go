package main

func maxProfit(prices []int) int {
	maxProfit := 0
	left := 0
	right := 1

	for right < len(prices) {
		if prices[left] > prices[right] {
			left = right
		} else {
			maxProfit = max(maxProfit, prices[right]-prices[left])
		}
		right++
	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
