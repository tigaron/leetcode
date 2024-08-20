package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i := range nums {
		diff := target - nums[i]
		if val, ok := numMap[diff]; ok {
			return []int{val, i}
		}
		numMap[nums[i]] = i
	}

	return nil
}
