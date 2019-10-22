package main

func uniqueMorseRepresentations(words []string) int {
	morse := [26]string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

	uniqueWords:= make(map[string]bool)

	for _, word := range words {
		morseWord := ""
		for _, litter := range word {
			morseWord += morse[litter - 'a']
		}
		uniqueWords[morseWord] = true
	}

	return len(uniqueWords)
}
