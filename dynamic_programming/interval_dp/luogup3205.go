package interval_dp

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 1001

var nums [MAXN]int
var n int
var MOD = 19650827

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		fmt.Sscanf(input, "%d", &n)

		for i := 1; i <= n; i++ {
			scanner.Scan()
			input = scanner.Text()

			fmt.Sscanf(input, "%d", &nums[i])
		}
	}
	if n == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(compute1(nums[:]))
	}
}

// 时间复杂度O(n^2)
// 严格位置依赖的动态规划
func compute1(nums []int) int {
	// 人的编号范围 : 1...n
	// dp[l][r][0] : 形成l...r的状况的方法数，同时要求l位置的数字是最后出现的
	// dp[l][r][1] : 形成l...r的状况的方法数，同时要求r位置的数字是最后出现的
	dp := make([][][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j < len(dp[i]); i++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 1; i < n; i++ {
		if nums[i] < nums[i+1] {
			dp[i][i+1][0] = 1
			dp[i][i+1][1] = 1
		}
	}
	for l := n - 2; l >= 1; l-- {
		for r := l + 2; r <= n; r++ {
			if nums[l] < nums[l+1] {
				dp[l][r][0] = (dp[l][r][0] + dp[l+1][r][0]) % MOD
			}
			if nums[l] < nums[r] {
				dp[l][r][0] = (dp[l][r][0] + dp[l+1][r][1]) % MOD
			}
			if nums[r] > nums[l] {
				dp[l][r][1] = (dp[l][r][1] + dp[l][r-1][0]) % MOD
			}
			if nums[r] > nums[r-1] {
				dp[l][r][1] = (dp[l][r][1] + dp[l][r-1][1]) % MOD
			}
		}
	}

	return (dp[1][n][0] + dp[1][n][1]) % MOD
}
