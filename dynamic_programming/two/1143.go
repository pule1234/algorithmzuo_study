package two

func longestCommonSubsequence(text1 string, text2 string) int {
	s1 := []byte(text1)
	s2 := []byte(text2)
	n := len(s1)
	m := len(s2)

	return f(s1, s2, n-1, m-1)
}

func f(s1, s2 []byte, i1, i2 int) int {
	if i1 < 0 || i2 < 0 {
		return 0
	}

	p1 := f(s1, s2, i1-1, i2-1)
	p2 := f(s1, s2, i1, i2-1)
	p3 := f(s1, s2, i1-1, i2)
	var p4 int

	if s1[i1] == s2[i2] {
		p4 = p1 + 1
	} else {
		p4 = 0
	}
	return max(max(p1, p2), max(p3, p4))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 为了避免很多边界讨论
// 很多时候往往不用下标来定义尝试，而是用长度来定义尝试
// 因为长度最短是0，而下标越界的话讨论起来就比较麻烦
func longestCommonSubsequence2(text1 string, text2 string) int {
	s1 := []byte(text1)
	s2 := []byte(text2)
	n := len(s1)
	m := len(s2)
}

func f2(s1, s2 []byte, len1, len2 int) int {
	if len1 == 0 || len2 == 0 {
		return 0
	}

	var ans int
	if s1[len1-1] == s2[len2-1] {
		ans = f2(s1, s2, len1-1, len2-1) + 1
	} else {
		ans = max(f2(s1, s2, len1-1, len2), f2(s1, s2, len1, len2-1))
	}

	return ans
}

// 记忆化搜索
func longestCommonSubsequence3(text1 string, text2 string) int {
	s1 := []byte(text1)
	s2 := []byte(text2)
	n := len(s1)
	m := len(s2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = -1
		}
	}

	return f3(s1, s2, n, m, dp)
}

func f3(s1, s2 []byte, len1, len2 int, dp [][]int) int {
	if len1 == 0 || len2 == 0 {
		return 0
	}

	if dp[len1][len2] != -1 {
		return dp[len1][len2]
	}

	var ans int
	if s1[len1-1] == s2[len2-1] {
		ans = f3(s1, s2, len1-1, len2-1, dp) + 1
	} else {
		ans = max(f3(s1, s2, len1, len2-1, dp), f3(s1, s2, len1-1, len2, dp))
	}
	dp[len1][len2] = ans

	return ans
}

// 动态规划
func longestCommonSubsequence4(text1 string, text2 string) int {
	s1 := []byte(text1)
	s2 := []byte(text2)
	n := len(s1)
	m := len(s2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for len1 := 1; len1 <= n; len1++ {
		for len2 := 1; len2 <= m; len2++ {
			if s1[len1-1] == s2[len2-1] {
				dp[len1][len2] = dp[len1-1][len2-1] + 1
			} else {
				dp[len1][len2] = max(dp[len1][len2-1], dp[len1-1][len2])
			}
		}
	}

	return dp[n][m]
}
