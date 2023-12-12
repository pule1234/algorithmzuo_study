package two

import (
	"math"
)

// 暴力递归
func minPathSum(grid [][]int) int {
	return f1(grid, len(grid), len(grid[0]))
}
func f1(grid [][]int, i, j int) int {
	if i == 0 && j == 0 {
		return grid[i][j]
	}

	up := math.MaxInt32
	left := math.MaxInt32

	if i-1 >= 0 {
		up = f1(grid, i-1, j)
	}
	if j-1 >= 0 {
		left = f1(grid, i, j-1)
	}
	return grid[i][j] + min(up, left)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 记忆化搜索
func minPathSum2(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
	}

	return f2(grid, n-1, m-1, dp)

}

func f2(grid [][]int, i, j int, dp [][]int) int {
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	ans := 0

	if i == 0 && j == 0 {
		ans = grid[i][j]
	} else {
		up := math.MaxInt32
		left := math.MaxInt32
		if i-1 >= 0 {
			up = f2(grid, i-1, j, dp)
		}
		if j-1 >= 0 {
			left = f2(grid, i, j-1, dp)
		}
		ans = grid[i][j] + min(left, up)
	}
	dp[i][j] = ans

	return ans
}

// 动态规划
func minPathSum3(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[n-1][m-1]
}

// 空间压缩的方法
func minPathSum4(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	dp := make([]int, m)
	dp[0] = grid[0][0]

	//先将第一行的数据录入
	for i := 1; i < m; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i := 1; i < n; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < m; j++ {
			dp[j] = grid[i][j] + min(dp[j-1], dp[j])
		}
	}

	return dp[len(dp)-1]
}
