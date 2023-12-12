package main

// 递归
func isMatch(s string, p string) bool {
	str := []byte(s)
	ptr := []byte(p)

	return f1(str, ptr, 0, 0)
}

func f1(str, ptr []byte, i, j int) bool {
	if i == len(str) {
		if j == len(ptr) {
			return true
		} else {
			// '*' 可以匹配空串
			return ptr[j] == '*' && f1(str, ptr, i, j+1)
		}
	} else if j == len(ptr) {
		return false
	} else {
		if ptr[j] != '*' {
			return (str[i] == ptr[j] || ptr[j] == '?') && f1(str, ptr, i+1, j+1)
		} else {
			return f1(str, ptr, i+1, j) || f1(str, ptr, i, j+1)
		}
	}
}

func isMatch(s string, p string) bool {
	str := []byte(s)
	ptr := []byte(p)
	n := len(str)
	m := len(ptr)
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
			ans = ptr[j] == '*' && f2(str, ptr, i, j+1, dp)
		}
	} else if j == len(ptr) {
		ans = false
	} else {
		if ptr[j] != '*' {
			ans = (str[i] == ptr[j] || ptr[j] == '?') && f2(str, ptr, i+1, j+1, dp)
		} else {
			ans = f2(str, ptr, i+1, j, dp) || f2(str, ptr, i, j+1, dp)
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
	for j := m - 1; j >= 0 && p[j] == '*'; j-- {
		dp[n][j] = true
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if ptr[j] != '*' {
				dp[i][j] = (str[i] == ptr[j] || ptr[j] == '?') && dp[i+1][j+1]
			} else {
				dp[i][j] = dp[i+1][j] || dp[i][j+1]
			}
		}
	}

	return dp[0][0]
}
