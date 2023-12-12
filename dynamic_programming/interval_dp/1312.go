package interval_dp

// 记忆化搜索
func minInsertions(s string) int {
	str := []byte(s)
	n := len(str)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	return f(str, 0, n-1, dp)

}

func f(str []byte, l, r int, dp [][]int) int {
	if dp[l][r] != -1 {
		return dp[l][r]
	}
	ans := 0
	if l == r {
		ans = 0
	} else if l+1 == r {
		if str[l] == str[r] {
			ans = 0
		} else {
			ans = 1
		}
	} else {
		if str[l] == str[r] {
			ans = f(str, l+1, r-1, dp)
		} else {
			ans = min(f(str, l+1, r, dp), f(str, l, r-1, dp)) + 1
		}
	}

	dp[l][r] = ans

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minInsertions(s string) int {
	str := []byte(s)
	n := len(str)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for l := 0; l < n-1; l++ {
		if str[l] == str[l+1] {
			dp[l][l+1] = 0
		} else {
			dp[l][l+1] = 1
		}
	}
	for l := n - 3; l >= 0; l-- {
		for r := l + 2; r < n; r++ {
			if str[l] == str[r] {
				dp[l][r] = dp[l+1][r-1]
			} else {
				dp[l][r] = min(dp[l][r-1], dp[l+1][r]) + 1
			}
		}
	}
	return dp[0][n-1]
}
