package main

import (
	"fmt"
	"regexp"
)

func main() {
}

func isMatch(s string, p string) bool {
	r, _ := regexp.MatchString(fmt.Sprintf("^%s$", p), s)
	return r
}
