package main

func romanToInt(s string) int {
	symbol := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var ans int

	for i := len(s) - 1; i >= 0; i-- {
		curr := symbol[string(s[i])]
		if i-1 >= 0 {
			next := symbol[string(s[i-1])]
			if curr > next {
				ans += curr - next
				i--
				continue
			}
		}
		ans += curr
	}

	return ans
}
