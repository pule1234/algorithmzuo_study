package main

import (
	"bufio"
	"fmt"
	"os"
)

// 01背包(模版)
// 给定一个正数t，表示背包的容量
// 有m个货物，每个货物可以选择一次
// 每个货物有自己的体积costs[i]和价值values[i]
// 返回在不超过总容量的情况下，怎么挑选货物能达到价值最大
// 返回最大的价值
// 测试链接 : https://www.luogu.com.cn/problem/P1048

const MAXM = 101
const MAXT = 1001

var cost [MAXM]int
var values [MAXM]int

var t, n int

// 严格位置依赖的动态规划
// n个物品编号1~n，第i号物品的花费cost[i]、价值val[i]
// cost、val数组是全局变量，已经把数据读入了
func compute1() int {
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, t+1)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= t; j++ {
			//不要i号货物
			dp[i][j] = dp[i-1][j]
			//要i号货物
			if j-cost[i] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-cost[i]]+values[i])
			}
		}
	}
	return dp[n][t]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		fmt.Sscanf(input, "%d %d", &t, &n)

		for i := 1; i <= n; i++ {
			scanner.Scan()
			input = scanner.Text()
			fmt.Sscanf(input, "%d %d", &cost[i], &values[i])
		}

		fmt.Println(compute1())
	}
}
