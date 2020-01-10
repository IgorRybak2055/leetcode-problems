package main

import "strings"

func main() {
	// for travis CI
}

func firstUniqChar(s string) int {
	a := make(map[rune]int)
	for _, ch := range s {
		a[ch]++
	}

	for _, ch := range s {
		if a[ch] == 1 {
			return strings.Index(s, string(ch))
		}
	}
	return -1

}
