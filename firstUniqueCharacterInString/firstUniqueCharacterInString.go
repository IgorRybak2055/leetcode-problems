package firstUniqueCharacterInString

import "strings"

func firstUniqChar(s string) int {
	a := make(map[rune]int, 0)
	for _, ch := range s{
		a[ch]++
	}

	for _, ch := range s {
		if a[ch] == 1 {
			return strings.Index(s, string(ch))
		}
	}
	return -1

}
