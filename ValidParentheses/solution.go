package main

func main() {

}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var (
		tmp = make([]rune,0, len(s)/2)
		m   = map[rune]rune{
			'(':')',
			'[':']',
			'{':'}',
		}
	)

	for _, val := range s {
		if len(tmp) == 0 {
			tmp = append(tmp, val)
			continue
		}
		if m[tmp[len(tmp) - 1]] == val {
			tmp = tmp[:len(tmp) - 1]
		} else {
			tmp = append(tmp, val)
		}
	}
	return len(tmp) == 0
}
