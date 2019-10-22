package ___maxIncreaseToKeepCitySkyline

func maxIncreaseKeepingSkyline(grid [][]int) int {
	maxRow, maxCol := make(map[int]int), make(map[int]int)

	for i := 0; i < len(grid); i++ {
		maxRow[i] = max(grid[i])
		for j := 0; j < len(grid[i]); j++ {
			if maxCol[j] < grid[i][j] {
				maxCol[j] = grid[i][j]
			}
		}
	}

	return culcResult(grid, maxRow, maxCol)
}

func culcResult(grid [][]int, maxRow map[int]int, maxCol map[int]int) int{
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			result += min(maxRow[i], maxCol[j]) - grid[i][j]
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {return a}
	return b
}

func max(nums []int) int {
	m := 0
	for _, val := range(nums){
		if val > m{
			m = val
		}
	}
	return m
}
