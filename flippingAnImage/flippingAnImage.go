package flippingAnImage

func flipAndInvertImage(A [][]int) [][]int {
	A = flip(A)
	return invert(A)
}

func flip(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		A[i] = reverse(A[i])
	}
	return A
}

func reverse(sl []int) []int {
	if len(sl) == 0 {
		return sl
	}
	return append(reverse(sl[1:]), sl[0])
}

func invert(A [][]int) [][]int {
	for i := 0; i < len(A); i++{
		for j := 0; j < len(A[i]); j++{
			switch A[i][j] {
			case 0:
				A[i][j] = 1
			case 1:
				A[i][j] = 0
			}
		}
	}
	return A
}

