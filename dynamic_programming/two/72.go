package two

func minDistance(word1 string, word2 string) int {
	return editDistance(word1, word2, 1, 1, 1)
}
func editDistance(word1 string, word2 string, a, b, c int) int {
	w1 := []byte(word1)
	w2 := []byte(word2)
	n := len(w1)
	m := len(w2)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = i * b
	}

	for j := 0; j <= m; j++ {
		dp[0][j] = j * a
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if w1[i-1] == w2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j]+b, dp[i][j-1]+a), dp[i-1][j-1]+c)
			}
		}
	}

	return dp[n][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
