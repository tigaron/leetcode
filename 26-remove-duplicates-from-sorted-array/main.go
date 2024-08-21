package main

func removeDuplicates(nums []int) int {
	prev := nums[0]
	ans := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[ans] = nums[i]
			ans++
		}

		prev = nums[i]
	}

	return ans
}
