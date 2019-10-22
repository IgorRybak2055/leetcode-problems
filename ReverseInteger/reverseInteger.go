package main

import (
	"fmt"
	"math"
	"strconv"
)

func main()  {
	// for travis CI
}

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = x * -1
	}

	ret := Reverse(strconv.Itoa(x)) * sign


	if ret> math.MaxInt32|| ret<math.MinInt32{
		return 0
	}
	return ret
}


func Reverse(s string) int {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	ret, err := strconv.Atoi(string(runes))
	if err != nil{
	 	fmt.Printf("Unexpected error: %#v", err.Error())
	}
	return ret
}
