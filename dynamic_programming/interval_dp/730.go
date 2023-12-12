package interval_dp

func countPalindromicSubsequences(s string) int {
	mod := 1000000007
	str := []byte(s)
	n := len(str)
	last := make([]int, 256)
	// left[i] : i位置的左边和s[i]字符相等且最近的位置在哪，不存在就是-1
	left := make([]int, n)
	for i := 0; i < len(last); i++ {
		last[i] = -1
	}

	for i := 0; i < n; i++ {
		left[i] = last[str[i]]
		last[s[i]] = i
	}
	// right[i] : i位置的右边和s[i]字符相等且最近的位置在哪，不存在就是n
	right := make([]int, n)
	for i := 0; i < len(last); i++ {
		last[i] = n
	}
	for i := n - 1; i >= 0; i-- {
		right[i] = last[str[i]]
		last[str[i]] = i
	}
	// dp[i][j] : i...j范围上有多少不同的回文子序列
	// 如果i>j，那么认为是无效范围dp[i][j] = 0
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if str[i] != str[j] {
				dp[i][j] = dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1] + mod
			} else { // s[i] == str[j]
				l := right[i]
				r := left[j]
				if l > r {
					//中间不存在相等的值
					// i...j的内部没有s[i]字符
					// a....a
					// i    j
					// (i+1..j-1) + a(i+1..j-1)a + a + aa
					dp[i][j] = dp[i+1][j-1]*2 + 2
				} else if l == r { // 中间有一个相等的值
					// i...j的内部有一个s[i]字符
					// a.....a......a
					// i     lr     j
					// (i+1..j-1) + a(i+1..j-1)a + aa
					dp[i][j] = dp[i+1][j-1]*2 + 1
				} else {
					// i...j的内部不只一个s[i]字符
					// a...a....这内部可能还有a但是不重要....a...a
					// i   l                             r   j
					// 因为要取模，所以只要发生减操作就+mod，讲解041同余原理
					dp[i][j] = dp[i+1][j-1]*2 - dp[l+1][r-1] + mod
				}
			}
			dp[i][j] %= mod
		}
	}
	return dp[0][n-1]
}
