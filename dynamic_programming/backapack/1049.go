package main

// 找到分割石头的方法，是两堆石头的差值更小
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, num := range stones {
		sum += num
	}
	// nums中随意选择数字
	// 累加和一定要 <= sum / 2
	// 又尽量接近
	near := nea(stones, sum/2)
	return sum - near - near
}

func nea(nums []int, t int) int {
	dp := make([]int, t+1)
	for _, num := range nums {
		for j := t; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return dp[t]
}
