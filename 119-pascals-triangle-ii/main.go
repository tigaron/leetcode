package main

import "math/big"

func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func getRow(rowIndex int) []int {
	var res []int

	rowFact := factorial(int64(rowIndex)) /* n! */

	for i := 0; i <= rowIndex; i++ { /* nCk = n! / (k! * (n-k)!) */
		iFact := factorial(int64(i))                    /* k! */
		rowMinusIFact := factorial(int64(rowIndex - i)) /* (n-k)! */

		denominator := new(big.Int).Mul(iFact, rowMinusIFact) /* k! * (n-k)! */
		fact := new(big.Int).Div(rowFact, denominator)        /* n! / (k! * (n-k)!) */

		factInt := int(fact.Int64())
		res = append(res, factInt)
	}

	return res
}
