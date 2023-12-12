package three

func knightProbability(n int, k int, row int, column int) float64 {
	dp := make([][][]float64, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]float64, n)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]float64, k+1)
			for z := 0; z < len(dp[i][j]); z++ {
				dp[i][j][z] = -1
			}
		}
	}

	return f(n, 0, 0, k, dp)
}

func f(n, i, j, k int, dp [][][]float64) float64 {
	if i < 0 || i >= n || j < 0 || j >= n {
		return 0
	}

	if dp[i][j][k] != -1 {
		return dp[i][j][k]
	}

	var ans float64
	if k == 0 {
		ans = 1
	} else {
		ans += (f(n, i-2, j+1, k-1, dp) / 8)
		ans += (f(n, i-1, j+2, k-1, dp) / 8)
		ans += (f(n, i+1, j+2, k-1, dp) / 8)
		ans += (f(n, i+2, j+1, k-1, dp) / 8)
		ans += (f(n, i+2, j-1, k-1, dp) / 8)
		ans += (f(n, i+1, j-2, k-1, dp) / 8)
		ans += (f(n, i-1, j-2, k-1, dp) / 8)
		ans += (f(n, i-2, j-1, k-1, dp) / 8)
	}

	dp[i][j][k] = ans
	return ans
}
