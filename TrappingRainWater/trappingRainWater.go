package main


func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}


func trap(height []int) int {
	var waterValue, leftBorder, rightBorder int
	for ind, val := range height {

		if val == 0 || ind < rightBorder {
			continue
		}
		leftBorder = ind
		if rightBorder = recessSearch(height, ind); rightBorder != -1 {
			minBorder := min(height[leftBorder], height[rightBorder])
			for i := leftBorder + 1; i < rightBorder; i++ {
				waterValue += minBorder - height[i]
			}
		}
	}

	return waterValue
}

func recessSearch(height []int, leftBorder int) int {
	ret := -1
	for i := leftBorder + 1; i < len(height); i++ {
		if height[i] >= height[leftBorder] {
			return i
		}
	}
	tmp := 0
	for i := leftBorder + 1; i < len(height); i++ {
		if tmp = max(tmp, height[i]); tmp == height[i] {
			ret = i
		}
	}

	return ret
}

