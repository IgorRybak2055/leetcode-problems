package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("234"))
}

func getLetters(button byte) (letters []string) {
	switch button {
	case '2':
		letters = []string{"a", "b", "c"}
	case '3':
		letters = []string{"d", "e", "f"}
	case '4':
		letters = []string{"g", "h", "i"}
	case '5':
		letters = []string{"j", "k", "l"}
	case '6':
		letters = []string{"m", "n", "o"}
	case '7':
		letters = []string{"p", "q", "r", "s"}
	case '8':
		letters = []string{"t", "u", "v"}
	case '9':
		letters = []string{"w", "x", "y", "z"}
	default:
		letters = []string{}
	}
	return
}

func merge(firstButtonLetters, secButtonLetters []string) (result []string) {
	for _, letterFromFirstButton := range firstButtonLetters {
		for _, letterFromSecButton := range secButtonLetters {
			result = append(result, letterFromFirstButton+letterFromSecButton)
		}
	}
	return
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	if len(digits) == 1 {
		return getLetters(digits[0])
	}

	var result = getLetters(digits[0])

	for i := 1; i < len(digits); i++ {
		result = merge(result, getLetters(digits[i]))
	}
	return result
}
