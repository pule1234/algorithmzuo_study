package two

func numDistinct(s string, t string) int {
	str := []byte(s)
	target := []byte(t)
	n := len(str)
	m := len(target)

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = 1
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i-1][j] //新来一个，默认为不相等
			if str[i-1] == target[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[n][m]
}
