package three

var mod = 1000000007

func numberOfPaths(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])
	return f(grid, n, m, k, 0, 0, 0)
}

func f(grid [][]int, n, m, k, i, j, r int) int {
	if i == n-1 && j == m-1 {
		if grid[i][j]%k == r {
			return 1
		}
		return 0
	}

	need := (k + r - (grid[i][j] % k)) % k
	ans := 0

	if i+1 < n {
		ans = f(grid, n, m, k, i+1, j, need)
	}
	if j+1 < m {
		ans = (ans + f(grid, n, m, k, i, j+1, need)) % mod
	}

	return ans
}

func numberOfPaths(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][][]int, n)
	for a := 0; a < n; a++ {
		dp[a] = make([][]int, m)
		for b := 0; b < m; b++ {
			dp[a][b] = make([]int, k)
			for c := 0; c < k; c++ {
				dp[a][b][c] = -1
			}
		}
	}

	return f2(grid, n, m, k, 0, 0, 0, dp)
}

func f2(grid [][]int, n, m, k, i, j, r int, dp [][][]int) int {
	if i == n-1 && j == m-1 {
		if grid[i][j]%r == 0 {
			return 1
		}
		return 0
	}

	if dp[i][j][r] != -1 {
		return dp[i][j][r]
	}

	need := (k + r - (grid[i][j] % k)) % k
	ans := 0
	if i+1 < n {
		ans = f2(grid, n, m, k, i+1, j, need, dp)
	}
	if j+1 < m {
		ans = (ans + f2(grid, n, m, k, i, j+1, need, dp)) % mod
	}

	dp[i][j][r] = ans
	return ans
}
