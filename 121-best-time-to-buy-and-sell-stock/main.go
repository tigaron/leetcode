package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProfit := 0
	currentProfit := 0

	for i := 1; i < len(prices); i++ {
		currentProfit = max(0, currentProfit+prices[i]-prices[i-1])
		maxProfit = max(maxProfit, currentProfit)
	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
