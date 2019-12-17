package main

import "strings"

func main()  {
	// for travis CI
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	var(
		start = 0
		end = len(s)-1
	)

	for start < end {
		if !isAlpha(s[start]) {
			start++
			continue
		}
		if !isAlpha(s[end]) {
			end--
			continue
		}

		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}
	return true
}

func isAlpha(c uint8) bool {
	if c >= 97 && c <= 122 {
		return true
	}

	if c >= 48 && c <= 57 {
		return true
	}

	return false
}


