package main

func main() {
	// for travis CI
}

func longestCommonPrefix(strs []string) string {
	if strs == nil {
		return ""
	}
	var resp string

	resp = strs[0]
	for _, val := range strs {
		if val == "" || resp == "" {
			return ""
		}

		if len(resp) > len(val) {
			resp = resp[:len(val)]
		} else {
			val = val[:len(resp)]
		}

		for i := 0; i < len(resp); i++ {
			if resp[i] != val[i] {
				resp = val[:i]
				break
			}
		}
	}
	return resp
}
