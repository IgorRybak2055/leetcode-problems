package removeOutermostParentheses

func removeOuterParentheses(S string) string {
	count, begin := 0, 1
	result := ""

	for i, val := range S {
		if val == '(' {
			count++
		}
		if val == ')'{
			count--
		}

		if count == 0 {
			result += S[begin:i]
			begin = i + 2
		}
	}
	return result
}