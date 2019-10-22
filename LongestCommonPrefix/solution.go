package main

import "fmt"

func main() {
	//fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	//fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
	fmt.Println(longestCommonPrefix([]string{"c", "c"}))
}

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
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
