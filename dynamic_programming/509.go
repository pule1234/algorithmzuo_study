package dynamic_programming

func fib(n int) int {
	return f1(n)
}

// 暴力递归
func f1(i int) int {
	if i == 0 {
		return 0
	}

	if i == 1 {
		return 1
	}

	return f1(i-1) + f1(i-2)
}

// 记忆化搜索
func fib2(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = -1
	}

	return f2(n, dp)
}

func f2(i int, dp []int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}

	if dp[i] != -1 { // 证明i的值已经求出来了
		return dp[i]
	}

	ans := f2(i-1, dp) + f2(i-2, dp)
	dp[i] = ans
	return ans
}

// 动态规划
func fib3(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1

	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 节省空间   用有限的变量代替数组
func fib4(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	last, lastlast := 1, 0

	for i := 2; i <= n; i++ {
		ans := lastlast + last
		lastlast = last
		last = ans
	}

	return last
}
