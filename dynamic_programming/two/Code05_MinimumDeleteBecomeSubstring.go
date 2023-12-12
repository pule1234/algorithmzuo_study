package two

import "math"

// 删除至少几个字符可以变成另一个字符串的子串
// 给定两个字符串s1和s2
// 返回s1至少删除多少字符可以成为s2的子串
// 对数器验证

func way1(s1, s2 string) int {
	str1 := []byte(s1)
	str2 := []byte(s2)
	n := len(str1)
	m := len(str2)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = i
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + dp[i-1][j]
			}
		}
	}

	ans := math.MaxInt32
	for j := 0; j <= m; j++ {
		ans = min(ans, dp[n][j])
	}

	return ans
}
