package dynamic_programming

func findSubstringInWraproundString(s string) int {
	n := len(s)
	str := make([]int, n)
	dp := make([]int, 26)
	for i := 0; i < n; i++ {
		str[i] = int(s[i] - 'a')
	}

	dp[str[0]] = 1
	var length = 1
	for i := 1; i < n; i++ {
		cur := str[i]
		pre := str[i-1]
		if (pre == 25 && cur == 0) || pre+1 == cur {
			length++
		} else {
			length = 1
		}
		// 会重复遇到相同的字符
		dp[cur] = max(dp[cur], length)
	}

	ans := 0
	for i := 0; i < len(dp); i++ {
		ans += dp[i]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
