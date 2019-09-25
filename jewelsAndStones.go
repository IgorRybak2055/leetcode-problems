package leetcodeProblems

import "strings"

func numJewelsInStones(J string, S string) int {
	count := 0
	for i := 0; i < len(J); i++{
		count += strings.Count(S, string(J[i]))
	}

	return count
}
