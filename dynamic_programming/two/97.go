package two

// 建立一个二维表，dp【i】[j-1] // 表示长度为j+i-1的s3 是否可以由长度为i的s1 和长度为j-1的s2组成
func isInterleave(s1 string, s2 string, s3 string) bool {

	if len(s1)+len(s2) != len(s3) {
		return false
	}

	str1 := []byte(s1)
	str2 := []byte(s2)
	str3 := []byte(s3)

	n := len(str1)
	m := len(str2)

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if str1[i-1] != str3[i-1] {
			break
		}
		dp[i][0] = true
	}

	for i := 1; i <= m; i++ {
		if str2[i-1] != str3[i-1] {
			break
		}
		dp[0][i] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = (str1[i-1] == str3[j+i-1] && dp[i-1][j]) || (str2[j-1] == str3[j+i-1] && dp[i][j-1])
		}
	}

	return dp[n][m]
}
