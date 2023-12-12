package two

func longestPalindromeSubseq(s string) int {
	str := []byte(s)
	return f(str, 0, len(str)-1)
}

func f(str []byte, l, r int) int {
	if l == r {
		return 0
	}

	if l+1 == r {
		if str[l] == str[r] {
			return 2
		} else {
			return 1
		}
	}

	if str[l] == str[r] {
		return 2 + f(str, l+1, r-1)
	} else {
		return max(f(str, l, r-1), f(str, l+1, r))
	}

}

// 记忆化搜索
func longestPalindromeSubseq2(s string) int {
	str := []byte(s)
	n := len(str)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	return f2(str, 0, n-1, dp)
}

func f2(str []byte, l, r int, dp [][]int) int {
	if l == r {
		return 0
	}

	if l+1 == r {
		if str[l] == str[r] {
			return 2
		} else {
			return 1
		}
	}

	var ans int
	if str[l] == str[r] {
		ans = 2 + f2(str, l+1, r-1, dp)
	} else {
		ans = max(f2(str, l+1, r, dp), f2(str, l, r-1, dp))
	}
	dp[l][r] = ans

	return ans
}

// 动态规划
func longestPalindromeSubseq3(s string) int {
	str := []byte(s)
	n := len(str)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for l := n - 1; l >= 0; l-- {
		dp[l][l] = 1
		if l+1 < n {
			if str[l] == str[l+1] {
				dp[l][l+1] = 2
			} else {
				dp[l][l+1] = 1
			}
		}

		for r := l + 2; r < n; r++ {
			if str[l] == str[r] {
				dp[l][r] = 2 + dp[l+1][r-1]
			} else {
				dp[l][r] = max(dp[l+1][r], dp[l][r-1])
			}
		}
	}
	return dp[0][n-1]
}
