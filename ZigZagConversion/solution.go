package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(convert("ABCD", 6))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	resp := make([]string, numRows)
	var (
		cnt     int
		allRows = true
		answer string
	)
	for _, litter := range s {
		if cnt == 0 {
			allRows = true
		}
		if cnt == numRows - 1 {
			allRows = false
		}

		resp[cnt] += string(litter)

		if !allRows {
			cnt--
		} else {
			cnt++
		}
	}

	for _, str := range resp {
		answer += strings.Replace(str, " ", "", -1)
	}

	return answer
}