package main

func removeDuplicates(nums []int) int {
	hm := make(map[int]int)

	for i := range nums {
		if _, ok := hm[nums[i]]; !ok {
			nums[len(hm)] = nums[i]
			hm[nums[i]] = 1
		} else {
			hm[nums[i]] += 1
		}
	}

	return len(hm)
}
