package two

// 有效涂色问题
// 给定n、m两个参数
// 一共有n个格子，每个格子可以涂上一种颜色，颜色在m种里选
// 当涂满n个格子，并且m种颜色都使用了，叫一种有效方法
// 求一共有多少种有效的涂色方法
// 1 <= n, m <= 5000
// 结果比较大请 % 1000000007 之后返回
// 对数器验证

const MAXN = 5001

var mod = 1000000007

func way1(n, m int) int {
	var dp [MAXN][MAXN]int

	for i := 1; i <= n; i++ {
		dp[i][1] = m
	}

	for i := 2; i <= n; i++ {
		for j := 2; j <= m; j++ {
			dp[i][j] = dp[i-1][j] * j                  // j中颜料都用了，剩下一个格子，在j种颜料中随便选取一个
			dp[i][j] = dp[i-1][j-1]*(m-j+1) + dp[i][j] // 用了j-1种颜料，剩下一个格子在m-j+1种选一个颜料填充
		}
	}

	return dp[n][m]
}
