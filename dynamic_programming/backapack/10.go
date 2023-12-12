package main

// 递归
func isMatch(s string, p string) bool {
	str := []byte(s)
	ptr := []byte(p)
	return f(str, ptr, 0, 0)
}

// s[i....]能不能被p[j....]完全匹配出来
// p[j]这个字符，一定不是'*'
func f(str, ptr []byte, i, j int) bool {
	if i == len(str) {
		if j == len(ptr) {
			return true
		} else {
			return j+1 < len(ptr) && ptr[j+1] == '*' && f(str, ptr, i, j+2)
		}
	} else if j == len(ptr) {
		return false
	} else {
		// s有后缀
		// p有后缀
		if j+1 == len(ptr) || ptr[j+1] != '*' {
			// s[i....]
			// p[j....]
			// 如果p[j+1]不是*，那么当前的字符必须能匹配：(s[i] == p[j] || p[j] == '?')
			// 同时，后续也必须匹配上：process1(s, p, i + 1, j + 1);
			return (str[i] == ptr[j] || ptr[j] == '.') && f(str, ptr, i+1, j+1)
		} else {
			// 如果p[j+1]是*
			// 完全背包！
			// s[i....]
			// p[j....]
			// 选择1: 当前p[j..j+1]是x*，就是不让它搞定s[i]，那么继续 : process1(s, p, i, j + 2)
			p1 := f(str, ptr, i, j+2)
			p2 := (str[i] == ptr[j] || ptr[j] == '.') && f(str, ptr, i+1, j)

			return p1 || p2
		}
	}
}

// 记忆化搜索
func isMatch(s string, p string) bool {
	str := []byte(s)
	ptr := []byte(p)
	n := len(str)
	m := len(ptr)
	// dp[i][j] == 0，表示没算过
	// dp[i][j] == 1，表示算过，答案是true
	// dp[i][j] == 2，表示算过，答案是false
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	return f2(str, ptr, 0, 0, dp)
}

func f2(str, ptr []byte, i, j int, dp [][]int) bool {
	if dp[i][j] != 0 {
		return dp[i][j] == 1
	}
	var ans bool
	if i == len(str) {
		if j == len(ptr) {
			ans = true
		} else {
			ans = j+1 < len(ptr) && ptr[j+1] == '*' && f2(str, ptr, i, j+2, dp)
		}
	} else if j == len(ptr) {
		ans = false
	} else {
		if j+1 == len(ptr) || ptr[j+1] != '*' {
			ans = (str[i] == ptr[j] || ptr[j] == '.') && f2(str, ptr, i+1, j+1, dp)
		} else {
			ans = f2(str, ptr, i, j+2, dp) || ((str[i] == ptr[j] || ptr[j] == '.') && f2(str, ptr, i+1, j, dp))
		}
	}
	if ans {
		dp[i][j] = 1
	} else {
		dp[i][j] = 2
	}

	return ans
}

func isMatch(s string, p string) bool {
	str := []byte(s)
	ptr := []byte(p)
	n := len(str)
	m := len(ptr)
	dp := make([][]bool, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[n][m] = true

	for j := m - 1; j >= 0; j-- {
		dp[n][j] = j+1 < m && ptr[j+1] == '*' && dp[n][j+2]
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if j+1 == m || ptr[j+1] != '*' {
				dp[i][j] = (str[i] == ptr[j] || ptr[j] == '.') && dp[i+1][j+1]
			} else {
				dp[i][j] = dp[i][j+2] || ((str[i] == ptr[j] || ptr[j] == '.') && dp[i+1][j])
			}
		}
	}
	return dp[0][0]
}
