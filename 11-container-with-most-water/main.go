package main

func maxArea(height []int) int {
	area := 0

	left := 0
	right := len(height) - 1

	for left < right {
		area = max(area, (right-left)*min(height[left], height[right]))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
