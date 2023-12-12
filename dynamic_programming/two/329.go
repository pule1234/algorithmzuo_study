package two

func longestIncreasingPath(matrix [][]int) int {
	ans := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			ans = max(ans, f(matrix, i, j))
		}
	}
}

func f(matrix [][]int, i, j int) int {
	next := 0
	if i > 0 && matrix[i][j] < matrix[i-1][j] {
		next = max(next, f(matrix, i-1, j))
	}

	if i+1 < len(matrix) && matrix[i][j] < matrix[i+1][j] {
		next = max(next, f(matrix, i+1, j))
	}

	if j > 0 && matrix[i][j] < matrix[i][j-1] {
		next = max(next, f(matrix, i, j-1))
	}
	if j+1 < len(matrix) && matrix[i][j] < matrix[i][j+1] {
		next = max(next, f(matrix, i, j+1))
	}

	return next + 1
}

func longestIncreasingPath(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	ans := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			ans = max(ans, f2(matrix, i, j, dp))
		}
	}
	return ans
}

func f2(matrix [][]int, i, j int, dp [][]int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}

	next := 0
	if i > 0 && matrix[i][j] < matrix[i-1][j] {
		next = max(next, f2(matrix, i-1, j, dp))
	}

	if i+1 < len(matrix) && matrix[i][j] < matrix[i+1][j] {
		next = max(next, f2(matrix, i+1, j, dp))
	}

	if j > 0 && matrix[i][j] < matrix[i][j-1] {
		next = max(next, f2(matrix, i, j-1, dp))
	}
	if j+1 < len(matrix[0]) && matrix[i][j] < matrix[i][j+1] {
		next = max(next, f2(matrix, i, j+1, dp))
	}

	dp[i][j] = next + 1
	return next + 1
}
