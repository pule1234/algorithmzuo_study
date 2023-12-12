package three

func isScramble(s1 string, s2 string) bool {
	str1 := []byte(s1)
	str2 := []byte(s2)
	n := len(str1)
	return f(str1, str2, 0, 0, n)
}

func f(s1, s2 []byte, l1, l2, length int) bool {
	if length == 1 {
		return s1[l1] == s2[l2]
	}

	for k := 1; k < length; k++ {
		if f(s1, s2, l1, l2, k) && f(s1, s2, l1+k, l2+k, length-k) {
			return true
		}
	}

	for i, j, k := l1+1, l2+length-1, 1; k < length; i, j, k = i+1, j-1, k+1 {
		if f(s1, s2, l1, j, k) && f(s1, s2, i, l2, length-k) {
			return true
		}
	}
	return false
}

func isScramble(s1 string, s2 string) bool {
	str1 := []byte(s1)
	str2 := []byte(s2)
	n := len(s1)

	// dp[l1][l2][len] : int 0 -> 没展开过
	// dp[l1][l2][len] : int -1 -> 展开过，返回的结果是false
	// dp[l1][l2][len] : int 1 -> 展开过，返回的结果是true
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	return f2(str1, str2, 0, 0, n, dp)
}

func f2(s1, s2 []byte, l1, l2, length int, dp [][][]int) bool {
	if length == 1 {
		return s1[l1] == s2[l2]
	}

	if dp[l1][l2][length] != 0 {
		return dp[l1][l2][length] == 1
	}

	ans := false
	for k := 1; k < length; k++ {
		if f2(s1, s2, l1, l2, k, dp) && f2(s1, s2, l1+k, l2+k, length-k, dp) {
			ans = true
			break
		}
	}
	if !ans {
		for i, j, k := l1+1, l2+length-1, 1; k < length; i, j, k = i+1, j-1, k+1 {
			if f2(s1, s2, l1, j, k, dp) && f2(s1, s2, i, l2, length-k, dp) {
				ans = true
				break
			}
		}
	}

	if ans {
		dp[l1][l2][length] = 1
	} else {
		dp[l1][l2][length] = -1
	}
	return ans
}
