package dynamic_programming

// 暴力尝试
func numDecodings1(s string) int {
	return f1([]byte(s), 0)
}

func f1(s []byte, i int) int {
	if i == len(s) {
		return 1
	}

	ans := 0

	if s[i] == '0' {
		return 0
	} else {
		ans = f1(s, i+1)
		if i+1 < len(s) && (s[i]-'0')*10+s[i+1]-'0' <= 26 {
			ans += f1(s, i+2)
		}
	}
	return ans
}

// 记忆化搜索
func numDecodings2(s string) int {
	str := []byte(s)
	dp := make([]int, len(str))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	return f2(str, 0, dp)
}

func f2(s []byte, i int, dp []int) int {
	if i == len(s) {
		return 1
	}

	if dp[i] != -1 {
		return dp[i]
	}

	ans := 0
	if s[i] == '0' {
		ans = 0
	} else {
		ans = f2(s, i+1, dp)
		if i+1 < len(s) && (s[i]-'0')*10+s[i+1]-'0' <= 26 {
			ans += f2(s, i+2, dp)
		}
	}

	dp[i] = ans

	return ans
}

// 动态规划
func numDecodings3(str string) int {
	s := []byte(str)
	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else {
			dp[i] = dp[i+1]
			if i+1 < len(s) && (s[i]-'0')*10+s[i+1]-'0' <= 26 {
				dp[i] += dp[i+2]
			}
		}
	}
	return dp[0]
}
