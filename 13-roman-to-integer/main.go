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

	for i := 0; i < len(s); i++ {
		curr := symbol[string(s[i])]
		next := 0
		if i+1 < len(s) {
			next = symbol[string(s[i+1])]
		}
		if curr < next {
			ans += next - curr
			i++
		} else {
			ans += curr
		}
	}

	return ans
}
