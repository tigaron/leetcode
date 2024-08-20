package main

func generate(numRows int) [][]int {
	results := make([][]int, numRows)
	for i := range numRows {
		results[i] = make([]int, i+1)
		results[i][0] = 1
	}

	if numRows == 1 {
		return results
	}

	for k := 1; k < numRows; k++ {
		for i := 1; i < len(results[k]); i++ {
			for j := i; j > 0; j-- {
				results[k][j] = results[k][j] + results[k][j-1]
			}
		}
	}

	return results
}
