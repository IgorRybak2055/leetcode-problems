package main

import "strings"

func main()  {
	// for travis CI
}

func numJewelsInStones(J string, S string) int {
	count := 0
	for i := 0; i < len(J); i++{
		count += strings.Count(S, string(J[i]))
	}

	return count
}
