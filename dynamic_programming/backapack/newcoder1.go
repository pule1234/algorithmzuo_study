package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 33001
const MAXM = 61

var cost [MAXM]int
var val [MAXM]int
var king [MAXM]bool
var fans [MAXM]int
var follows [MAXM][2]int
var dp [MAXN]int
var n, m int

func clean() {
	for i := 1; i <= m; i++ {
		fans[i] = 0
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		fmt.Sscanf(input, "%d %d", &n, &m)
		clean()
		for i := 1; i <= m; i++ {
			scanner.Scan()
			input = scanner.Text()
			var v, p, q int
			fmt.Sscanf(input, "%d %d %d", &v, &p, &q)
			cost[i] = v
			val[i] = v * p
			king[i] = q == 0
			if q != 0 {
				follows[q][fans[q]] = i
				fans[q]++
			}
		}
		fmt.Println(compute())
	}
}
func compute() int {
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	// p : 上次展开的主商品编号
	p := 0

	for i := 1; i <= m; i++ {
		if king[i] {
			for j := 0; j <= n; j++ {
				// dp[i][j] : 0...i范围上，只关心主商品，并且进行展开
				//            花费不超过j的情况下，获得的最大收益
				// 可能性1 : 不考虑当前主商品
				dp[i][j] = dp[p][j]
				if j-cost[i] >= 0 {
					// 可能性2 : 考虑当前主商品，只要主
					dp[i][j] = max(dp[i][j], dp[p][j-cost[i]]+val[i])
				}
				fan1 := -1
				fan2 := -1
				if fans[i] >= 1 {
					fan1 = follows[i][0]
				}
				if fans[i] >= 2 {
					fan2 = follows[i][1]
				}

				if fan1 != -1 && j-cost[i]-cost[fan1] >= 0 {
					// 可能性3 : 主 + 附1
					dp[i][j] = max(dp[i][j], dp[p][j-cost[i]-cost[fan1]]+val[i]+val[fan1])
				}
				//
				//if fan2 != -1 && j-cost[i]-cost[fan2] >= 0 {
				//	dp[i][j] = max(dp[i][j], dp[p][j-cost[i]-cost[fan2]]+val[i]+val[fan2])
				//}

				if fan2 != -1 && j-cost[i]-cost[fan2] >= 0 {
					// 可能性4 : 主 + 附2
					dp[i][j] = max(dp[i][j], dp[p][j-cost[i]-cost[fan2]]+val[i]+val[fan2])
				}

				if fan1 != -1 && fan2 != -1 && j-cost[i]-cost[fan1]-cost[fan2] >= 0 {
					// 可能性5 : 主 + 附1 + 附2
					dp[i][j] = max(dp[i][j], dp[p][j-cost[i]-cost[fan1]-cost[fan2]]+val[i]+val[fan1]+val[fan2])
				}
			}
			p = i
		}

	}

	return dp[p][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
