package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strStr("hello", "8"))
}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

