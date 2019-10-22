package main

import (
	"math"
)

func main()  {
	// for travis CI
}

func myAtoi(str string) int {
	var (
		result, sign = 0, 1
		start bool
	)
	for _, litter := range str {
		if litter == ' ' && !start {
			continue
		}
		if litter == '+' && !start {
			sign = 1
			start = true
			continue
		}
		if litter == '-' && !start {
			sign = -1
			start = true
			continue
		}

		if litter >= '0' && litter <= '9' {
			start = true
			result = result * 10 + (int(litter) - 48)
			if result * sign > math.MaxInt32 {
				return math.MaxInt32
			}
			if result * sign < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}
	return result * sign
}

