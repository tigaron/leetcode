package main

func strStr(haystack string, needle string) int {
	nh := len(haystack)
	nn := len(needle)

	if nh < nn {
		return -1
	}

	if haystack == needle {
		return 0
	}

	for i := 0; i <= nh-nn; i++ {
		if string(haystack[i:i+nn]) == needle {
			return i
		}
	}

	return -1
}
