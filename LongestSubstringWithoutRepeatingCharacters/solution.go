package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	var start, lenSubStr int

	for ind, letter := range s {
		if letterInd, ok := m[letter]; ok && letterInd >= start {
			start = letterInd + 1
			m[letter] = ind
			continue
		}

		if ind - start + 1 > lenSubStr {
			lenSubStr = ind - start + 1
		}

		m[letter] = ind
	}

	return lenSubStr
}
