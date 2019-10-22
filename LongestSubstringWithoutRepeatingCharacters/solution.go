package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	var start, max int
	for ind, ch := range s {
		if litterInd, ok := m[ch]; ok && m[ch] >= start {
			start = litterInd + 1
			m[ch] = ind
			continue
		}

		if length := ind - start + 1; length > max {
			max = length
		}
		m[ch] = ind
	}
	return max
}
