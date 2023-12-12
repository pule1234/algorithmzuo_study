package two

import "fmt"

// 节点数为n高度不大于m的二叉树个数
// 现在有n个节点，计算出有多少个不同结构的二叉树
// 满足节点个数为n且树的高度不超过m的方案
// 因为答案很大，所以答案需要模上1000000007后输出
// 测试链接 : https://www.nowcoder.com/practice/aaefe5896cce4204b276e213e725f3ea
// 请同学们务必参考如下代码中关于输入、输出的处理
// 这是输入输出处理效率很高的写法
// 提交以下所有代码，把主类名改成Main，可以直接通过
const MAXN = 51

var Mod = 1000000007

func way1(num int, level int) int {
	dp := make([][]int, MAXN)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, MAXN)
		for j := 0; j < MAXN; j++ {
			dp[i][j] = -1
		}
	}
	return computer1(num, level, dp)
}

func computer1(n, m int, dp [][]int) int {
	if n == 0 {
		return 1
	}

	if m == 0 {
		return 0
	}

	if dp[n][m] != -1 {
		return dp[n][m]
	}

	ans := 0
	for k := 0; k < n; k++ {
		ans = (ans + (computer1(k, m-1, dp)*computer1(n-k-1, m-1, dp))%Mod) % Mod
	}

	dp[n][m] = ans

	return ans
}

func main() {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		} else {
			fmt.Println(way1(a, b))
		}
	}
}

func way2(num, level int) int {
	dp := make([][]int, MAXN)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, MAXN)
	}
	for i := 0; i < level; i++ {
		dp[0][i] = 1
	}

	for i := 1; i <= num; i++ {
		for j := 1; j <= level; j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = (dp[i][j] + dp[k][j-1]*dp[i-k-1][j-1]%Mod) % Mod
			}
		}
	}

	return dp[num][level]
}
