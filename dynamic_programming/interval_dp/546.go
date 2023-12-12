package interval_dp

func removeBoxes(boxes []int) int {
	n := len(boxes)
	dp := make([][][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, n)
		}
	}
	return f(boxes, 0, n-1, 0, dp)
}

// boxes[l....r]范围上要去消除，前面跟着k个连续的和boxes[l]颜色一样的盒子
func f(boxes []int, l, r, k int, dp [][][]int) int {
	if l > r {
		return 0
	}

	if dp[l][r][k] > 0 {
		return dp[l][r][k]
	}
	s := l
	for s+1 <= r && boxes[l] == boxes[s+1] {
		s++
	}
	// boxes[l...s]都是一种颜色，boxes[s+1]就不是同一种颜色了
	cnt := k + s - l + 1
	// 可能性1 : 前缀先消
	ans := cnt*cnt + f(boxes, s+1, r, 0, dp)
	// 可能性2 : 讨论前缀跟着哪个后，一起消掉
	for m := s + 2; m <= r; m++ {
		if boxes[l] == boxes[m] && boxes[m-1] != boxes[m] {
			// boxes[l] == boxes[m]是必须条件
			// boxes[m - 1] != boxes[m]是剪枝条件，避免不必要
			ans = max(ans, f(boxes, s+1, m-1, 0, dp)+f(boxes, m, r, cnt, dp))
		}
	}

	dp[l][r][k] = ans

	return ans
}
