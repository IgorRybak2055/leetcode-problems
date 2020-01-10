package main

import "fmt"

func main()  {
	fmt.Println(isPalindrome("race a car"))
	// fmt.Println(isPalindrome("Abba"))
}

func isPalindrome(s string) bool {
	var(
		start = 0
		end = len(s)-1
	)

	for start < end {
		if !(s[start] >= 97 && s[start] <= 122) && !(s[start] >= 48 && s[start] <= 57) && !(s[start] >= 65 && s[start] <= 90) {
			start++
			continue
		}

		if !(s[end] >= 97 && s[end] <= 122) && !(s[end] >= 48 && s[end] <= 57) && !(s[end] >= 65 && s[end] <= 90) {
			end--
			continue
		}

		if !((s[start] != s[end]) || (s[start] - 32 != s[end]) || (s[start] != s[end] - 32)) { // некорректно работает
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

