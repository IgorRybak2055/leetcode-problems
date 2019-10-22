package main

import "strings"

func main()  {
	// for travis CI
}

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}