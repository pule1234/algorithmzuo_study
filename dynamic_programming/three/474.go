package three

var zores, ones int

func findMaxForm(strs []string, m int, n int) int {
	return f1(strs, 0, m, n)
}

func f1(strs []string, i, m, n int) int {
	if i == len(strs) {
		return 0
	}

	// 不选择当前字符串
	p1 := f1(strs, i+1, m, n)

	// 选择当前位置
	p2 := 0
	zerosAndOnes(strs[i])
	if zores <= m && ones <= n {
		p2 = f1(strs, i, m-zores, n-ones) + 1
	}

	return max(p1, p2)
}

func zerosAndOnes(str string) {
	zores = 0
	ones = 0

	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			zores++
		} else {
			ones++
		}
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func findMaxForm2(strs []string, m int, n int) int {

	dp := make([][][]int, len(strs))
	for i := 0; i < len(strs); i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
			for k := 0; k <= n; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	return f2(strs, 0, m, n, dp)
}

func f2(strs []string, i, m, n int, dp [][][]int) int {
	if i == len(strs) {
		return 0
	}

	if dp[i][m][n] != -1 {
		return dp[i][m][n]
	}

	p1 := f2(strs, i+1, m, n, dp)
	p2 := 0
	zerosAndOnes(strs[i])
	if zores <= m && ones <= n {
		p2 = f2(strs, i+1, m-zores, n-ones, dp) + 1
	}

	ans := max(p1, p2)
	dp[i][m][n] = ans
	return ans
}

func findMaxForm3(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs)+1)
	for i := 0; i < len(strs); i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	length := len(strs)

	for i := length - 1; i >= 0; i-- {
		zerosAndOnes(strs[i])
		for z := 0; z <= m; z++ {
			for o := 0; o <= n; o++ {
				p1 := dp[i+1][z][o]
				p2 := 0
				if zores <= z && ones <= o {
					p2 = 1 + dp[i+1][z-zores][o-ones]
				}
				dp[i][z][o] = max(p1, p2)
			}
		}
	}
	return dp[0][m][n]
}
