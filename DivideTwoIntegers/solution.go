package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(7, 0))
}

func divide(dividend int, divisor int) int {
	var sign = 1

	if divisor == 0 {
		return 0
	}

	if dividend < 0 {
		sign *= -1
	}

	if divisor < 0 {
		sign *= -1
	}

	dvd := math.Abs(float64(dividend))
	dvs := math.Abs(float64(divisor))

	if dvs == 1 {
		if int(dvd) * sign > math.MaxInt32 {
			return math.MaxInt32
		}
		if int(dvd) * sign < math.MinInt32 {
			return math.MinInt32
		}

		fl := int(dvs) == divisor

		if fl {
			return dividend
		}

		return int(dvd) * sign
	}

	quotient := 0
	for dvd >= dvs {
		dvd -= dvs
		quotient++
	}

	if sign * quotient > math.MaxInt32 {
		return math.MaxInt32
	}

	if sign * quotient < math.MinInt32 {
		return math.MinInt32
	}

	return sign * quotient
}

