package main

func main()  {
	// for travis CI
}

func numSquares(n int) int {
	qtyNums := make([]int, n+1)
	qtyNums[0] = 0

	for i := 1; i<=n;i++{
		qtyNums[i] = i
		for j := 1; j*j <= i;j++ {
			qtyNums[i] = Min(qtyNums[i], qtyNums[i - j*j] + 1)
		}
	}

	return qtyNums[n]
}


func Min(a,b int) int {
	if a < b {return a}
	return b
}
