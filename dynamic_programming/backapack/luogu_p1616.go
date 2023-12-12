package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXM = 1001
const MAXT = 10000001

var cost [MAXM]int
var val [MAXM]int
var dp [MAXT]int
var t, m int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		fmt.Sscanf(input, "%d %d", &t, &m)

		for i := 1; i <= m; i++ {
			scanner.Scan()
			input = scanner.Text()

			fmt.Sscanf(input, "%d %d", &cost[i], &val[i])
		}
	}

	fmt.Println(compute2())
}

func compute() int {
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, t+1)
	}

	for i := 1; i <= m; i++ {
		for j := 0; j <= t; j++ {
			dp[i][j] = dp[i-1][j]
			if j-cost[i] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-cost[i]]+val[i])
			}
		}
	}
	return dp[m][t]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func compute2() int {
	for i := 1; i <= t; i++ {
		dp[i] = 0
	}
	for i := 1; i <= m; i++ {
		for j := cost[i]; j <= t; j++ {
			dp[j] = max(dp[j], dp[j-cost[i]]+val[i])
		}
	}
	return dp[t]
}
