package main

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func countOfAtoms(formula string) string {
	stack := []map[string]int{}
	stack = append(stack, map[string]int{})

	i := 0
	n := len(formula)

	for i < n {
		if formula[i] == '(' {
			stack = append(stack, map[string]int{})
			i++
		} else if formula[i] == ')' {
			i++
			start := i
			for i < n && unicode.IsDigit(rune(formula[i])) {
				i++
			}
			multiplicity, _ := strconv.Atoi(formula[start:i])
			if multiplicity == 0 {
				multiplicity = 1
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for elem, count := range top {
				stack[len(stack)-1][elem] += count * multiplicity
			}
		} else {
			start := i
			i++
			for i < n && unicode.IsLower(rune(formula[i])) {
				i++
			}
			element := formula[start:i]

			start = i
			for i < n && unicode.IsDigit(rune(formula[i])) {
				i++
			}
			count, _ := strconv.Atoi(formula[start:i])
			if count == 0 {
				count = 1
			}

			stack[len(stack)-1][element] += count
		}
	}

	result := stack[0]
	elements := make([]string, 0, len(result))
	for elem := range result {
		elements = append(elements, elem)
	}
	sort.Strings(elements)

	var sb strings.Builder
	for _, elem := range elements {
		sb.WriteString(elem)
		if result[elem] > 1 {
			sb.WriteString(strconv.Itoa(result[elem]))
		}
	}

	return sb.String()
}
