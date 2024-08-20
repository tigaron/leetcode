package main

func isValid(s string) bool {
	stack := [][]string{}

	i := 0
	n := len(s)

	for i < n {
		switch string(s[i]) {
		case "(":
			stack = append(stack, []string{"("})
			i++
		case ")":
			if len(stack) > 0 {
				if stack[len(stack)-1][0] != "(" {
					return false
				}
				stack = stack[:len(stack)-1]
				i++
			} else {
				return false
			}
		case "[":
			stack = append(stack, []string{"["})
			i++
		case "]":
			if len(stack) > 0 {
				if stack[len(stack)-1][0] != "[" {
					return false
				}
				stack = stack[:len(stack)-1]
				i++
			} else {
				return false
			}
		case "{":
			stack = append(stack, []string{"{"})
			i++
		case "}":
			if len(stack) > 0 {
				if stack[len(stack)-1][0] != "{" {
					return false
				}
				stack = stack[:len(stack)-1]
				i++
			} else {
				return false
			}
		default:
			return false
		}
	}

	return len(stack) == 0
}
