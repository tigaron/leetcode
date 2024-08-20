package main

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	low, high := 0, nums[len(nums)-1]-nums[0]

	for low < high {
		mid := low + (high-low)/2

		count, start := 0, 0

		for end := 0; end < len(nums); end++ {
			for nums[end]-nums[start] > mid {
				start++
			}
			count += end - start
		}

		if count >= k {
			high = mid
		} else {
			low = mid + 1
		}

	}

	return low
}
