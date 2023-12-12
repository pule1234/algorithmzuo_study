package dynamic_programming

import "fmt"

func longestValidParentheses(s string) int {
	str := []byte(s)
	n := len(str)
	dp := make([]int, n)
	// dp[0...n-1]
	// dp[i] : 子串必须以i位置的字符结尾的情况下，往左整体有效的最大长度
	ans := 0

	for i := 1; i < n; i++ {
		if str[i] == ')' {
			p := i - dp[i-1] - 1
			if p >= 0 && str[p] == '(' {
				if p-1 >= 0 {
					dp[i] = dp[i-1] + dp[p-1] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		ans = max(ans, dp[i])
		fmt.Println(ans)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
