package main

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	left, right := 0, nums[len(nums)-1]-nums[0]

	for left < right {
		mid := left + (right-left)/2

		if countPairsWithDistanceLessThanOrEqual(nums, mid) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func countPairsWithDistanceLessThanOrEqual(arr []int, mid int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		j := i + 1
		for j < len(arr) && arr[j]-arr[i] <= mid {
			j++
		}
		count += j - i - 1
	}
	return count
}
