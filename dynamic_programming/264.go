package dynamic_programming

// log(n) n代表第n个丑数
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	for i, i2, i3, i5 := 2, 1, 1, 1; i <= n; i++ {
		a := dp[i2] * 2
		b := dp[i3] * 3
		c := dp[i5] * 5

		cur := min(min(a, b), c)
		if cur == a {
			i2++
		}
		if cur == b {
			i3++
		}
		if cur == c {
			i5++
		}
		dp[i] = cur
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
