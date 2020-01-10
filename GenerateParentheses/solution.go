package main

import "fmt"

func main() {
	a := generateParenthesis(2)

	fmt.Println(len(a), a)
}

func generateParenthesis(n int) []string {
	answer := make([]string, 0)
	var sign []byte
	backtrack(&answer, sign, 0, 0, n)
	return answer
}

func backtrack(answer *[]string, sign []byte, open, close, max int) {
	if len(sign) == max*2 {
		*answer = append(*answer, string(sign))
		return
	}
	if open < max {
		backtrack(answer, append(sign, '('), open+1, close, max)
	}
	if close < open {
		backtrack(answer, append(sign, ')'), open, close+1, max)
	}
}