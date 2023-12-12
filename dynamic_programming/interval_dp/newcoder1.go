package interval_dp

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// 完成配对需要的最少字符数量
// 给定一个由'['、']'、'('，')'组成的字符串
// 请问最少插入多少个括号就能使这个字符串的所有括号正确配对
// 例如当前串是 "([[])"，那么插入一个']'即可满足
// 输出最少需要插入多少个字符
// 测试链接 : https://www.nowcoder.com/practice/e391767d80d942d29e6095a935a5b96b

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	var str string
	fmt.Sscanf(input, "%s", &str)
	fmt.Println(compute(str))
}
func compute(str string) int {
	s := []byte(str)
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	return f(s, 0, n-1, dp)
}

// 让s[l...r]配对至少需要几个字符
func f(s []byte, l, r int, dp [][]int) int {
	if l == r {
		return 1
	}
	if l == r-1 {
		if (s[l] == '(' && s[r] == ')') || (s[l] == '[' && s[r] == ']') {
			return 0
		}
		return 2
	}

	if dp[l][r] != -1 {
		return dp[l][r]
	}
	p1 := math.MaxInt32
	if (s[l] == '(' && s[r] == ')') || (s[l] == '[' && s[r] == ']') {
		p1 = f(s, l+1, r-1, dp)
	}
	p2 := math.MaxInt32
	for m := l; m < r; m++ {
		p2 = min(p2, f(s, l, m, dp)+f(s, m+1, r, dp))
	}

	ans := min(p1, p2)
	dp[l][r] = ans
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
