package main

import (
	"log"
	"strings"
)

func main() {
	var (
		s     = "barfoothefoobarman"
		words = []string{"foo", "bar"}
	)

	log.Println(findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {
	var (
		result             []int
		left, right, match int
	)

	if s == "" || words == nil || len(s) < len(strings.Join(words, "")) {
		return result
	}

	wordLen := len(words[0])

	appear := make(map[string]int)
	for _, word := range words {
		appear[word] = appear[word] + 1
	}

	dic := appear

	for left <= len(s)-len(words)*wordLen {
		if _, ok := dic[s[left:left+wordLen]]; ok {
			right = left
			for val, ok := dic[s[right:right+wordLen]]; val > 0 && ok; {
				dic[s[right:right+wordLen]] -= 1
				if dic[s[right:right+wordLen]] == 0 {
					match += 1
				}
				right += wordLen

				if match == len(dic) {
					result = append(result, left)
				}
			}
			dic = appear
		}
		left++
	}
	return result
}