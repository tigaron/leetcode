package main

func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true // Base case: empty array is a valid partition

	for i := 2; i <= n; i++ {
		if i >= 2 && nums[i-1] == nums[i-2] {
			dp[i] = dp[i] || dp[i-2] // Case where last two numbers are equal
		}
		if i >= 3 && nums[i-1] == nums[i-2] && nums[i-2] == nums[i-3] {
			dp[i] = dp[i] || dp[i-3] // Case where last three numbers are equal
		}
		if i >= 3 && nums[i-1]-nums[i-2] == 1 && nums[i-2]-nums[i-3] == 1 {
			dp[i] = dp[i] || dp[i-3] // Case where last three numbers are consecutive
		}
	}

	return dp[n]
}
