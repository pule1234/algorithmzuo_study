package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MAXN = 1001

// arr[i][0] i号物品的体积
// arr[i][1] i号物品的价值
// arr[i][2] i号物品的组号
var arr [MAXN][3]int

var m, n int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}

		fmt.Sscanf(input, "%d %d", &m, &n)

		for i := 1; i <= n; i++ {
			scanner.Scan()
			input = scanner.Text()
			fmt.Sscanf(input, "%d %d %d", &arr[i][0], &arr[i][1], &arr[i][2])
		}

		sort.Slice(arr, func(i, j int) bool {
			return arr[i][2] < arr[j][2]
		})

		fmt.Println(compute())
	}
}

func compute() int {
	teams := 1
	for i := 2; i <= n; i++ {
		if arr[i-1][2] != arr[i][2] {
			teams++
		}
	}
	// 组的编号1~teams
	// dp[i][j] : 1~i是组的范围，每个组的物品挑一件，容量不超过j的情况下，最大收益
	dp := make([][]int, teams+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	// dp[0][....] = 0
	for start, end, i := 1, 2, 1; start <= n; i++ {
		for end <= n && arr[end][2] == arr[start][2] {
			end++
		}
		// start ... end-1 -> i组
		for j := 0; j <= m; j++ {
			// arr[start...end-1]是当前组，组号一样
			// 其中的每一件商品枚举一遍
			dp[i][j] = dp[i-1][j]
			for k := start; k < end; k++ {
				if j-arr[k][0] >= 0 {
					dp[i][j] = max(dp[i][j], dp[i-1][j-arr[k][0]]+arr[k][1])
				}
			}
		}
		start = end
		end++
	}

	return dp[teams][m]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
