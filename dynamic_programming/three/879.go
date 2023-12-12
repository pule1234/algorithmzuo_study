package three

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	return f1(group, profit, 0, n, minProfit)
}

// i : 来到i号工作
// r : 员工额度还有r人，如果r<=0说明已经没法再选择工作了
// s : 利润还有s才能达标，如果s<=0说明之前的选择已经让利润达标
func f1(g, p []int, i, r, s int) int {
	if r <= 0 {
		if s <= 0 {
			return 1
		}
		return 0
	}

	if i == len(g) {
		if s <= 0 {
			return 1
		}
		return 0
	}

	// 不要当前工作
	p1 := f1(g, p, i+1, r, s)

	//要当前工作
	p2 := 0
	if g[i] <= r {
		p2 = f1(g, p, i+1, r-g[i], s-p[i])
	}

	return p1 + p2
}

var mod = 1000000007

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	m := len(group)
	dp := make([][][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, minProfit+1)
			for k := 0; k < len(dp[i][j]); k++ {
				dp[i][j][k] = -1
			}
		}
	}
	return f2(group, profit, 0, n, minProfit, dp)
}

func f2(g, p []int, i, r, s int, dp [][][]int) int {
	if r <= 0 {
		if s == 0 {
			return 1
		}
		return 0
	}

	if i == len(g) {
		if s == 0 {
			return 1
		}
		return 0
	}

	if dp[i][r][s] != -1 {
		return dp[i][r][s]
	}

	p1 := f2(g, p, i+1, r, s, dp)
	p2 := 0

	if g[i] <= r {
		p2 = f2(g, p, i+1, r-g[i], max(0, s-p[i]), dp)
	}

	ans := (p1 + p2) % mod
	dp[i][r][s] = ans
	return ans
}

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	m := len(group)
	dp := make([][][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, minProfit+1)
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			for z := minProfit; z >= 0; z-- {
				p1 := dp[i+1][j][z]
				p2 := 0

				if group[i] <= j {
					p2 = dp[i+1][j-group[i]][max(0, z-profit[i])]
				}

				dp[i][j][z] = p1 + p2
			}
		}
	}

	return dp[0][n][minProfit]
}
