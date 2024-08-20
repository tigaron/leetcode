package main

func generate(numRows int) [][]int {
	results := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		results[i] = make([]int, i+1)
		results[i][0], results[i][i] = 1, 1

		for j := 1; j < i; j++ {
			results[i][j] = results[i-1][j-1] + results[i-1][j]
		}
	}

	return results
}
