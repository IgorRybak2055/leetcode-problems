package main

func main()  {
	// for travis CI
}
func longestPalindrome(s string) string {
	var (
		start, end int
		strLen = len(s)
	)

	for i := 0; i < strLen; i++ {
		// find longest palindromic substring with odd length and center at i
		leftInd := i
		rightInd := i
		for leftInd >= 0 && rightInd < strLen && s[leftInd] == s[rightInd] {
			leftInd--
			rightInd++
		}
		if rightInd-leftInd-1 > end-start {
			start = leftInd + 1
			end = rightInd
		}
		
		leftInd, rightInd = i, i + 1
		for leftInd >= 0 && rightInd < strLen && s[leftInd] == s[rightInd] {
			leftInd--
			rightInd++
		}
		if rightInd-leftInd-1 > end-start {
			start = leftInd + 1
			end = rightInd
		}
	}

	return s[start:end]
}

