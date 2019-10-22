package main

func main()  {
	
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var reverseNum int
	num := x

	for x > 0 {
		reverseNum = reverseNum * 10 + x%10
		x = x/10
	}

	return reverseNum == num
}