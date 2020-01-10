package main

import (
	"log"
	"strconv"
)

func main() {
	log.Println(countAndSay(4))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := countAndSay(n - 1)

	result, length := "", 1
	for i := 0; i < len(prev)-1; i++ {
		if prev[i] != prev[i+1] {
			result = result + strconv.Itoa(length) + string(prev[i])
			length = 1
		} else {
			length++
		}
	}

	return result + strconv.Itoa(length) + string(prev[len(prev)-1])
}
