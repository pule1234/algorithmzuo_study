package interval_dp

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func strangePrinter(s string) int {
	str := []byte(s)
	n := len(str)

	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	dp[n-1][n-1] = 1

	for i := 0; i < n-1; i++ {
		dp[i][i] = 1
		if s[i] == s[i+1] {
			dp[i][i+1] = 1
		} else {
			dp[i][i+1] = 2
		}
	}

	for l, ans := n-3, 0; l >= 0; l-- {
		for r := l + 2; r < n; r++ {
			if s[l] == s[r] {
				dp[l][r] = dp[l][r-1]
			} else {
				ans = math.MaxInt32
				for m := l; m < r; m++ {
					ans = min(ans, dp[l][m]+dp[m+1][r])
				}
				dp[l][r] = ans
			}
		}
	}

	return dp[0][n-1]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	var s string
	fmt.Sscanf(input, "%s", &s)
	fmt.Println(strangePrinter(s))
}
