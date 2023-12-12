package Dual_pointers

//func trap(height []int) int {
//	n := len(height)
//	// 左边的最大值
//	lmax := make([]int, n)
//	//右边的最大值
//	rmax := make([]int, n)
//
//	lmax[0] = height[0]
//	for i := 1; i < n; i++ {
//		lmax[i] = max(lmax[i-1], height[i])
//	}
//	rmax[n-1] = height[n-1]
//	for i := n - 2; i >= 0; i-- {
//		rmax[i] = max(rmax[i+1], height[i])
//	}
//
//	ans := 0
//	for i := 1; i < n-1; i++ { // 第一个和最后一个柱子不可以装水
//		ans += max(0, min(lmax[i-1], rmax[i+1])-height[i])
//	}
//	return ans
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 利用执行两边的操作
func trap(height []int) int {
	n := len(height)
	l, r := 1, n-2
	lmax, rmax := height[0], height[n-1]
	ans := 0

	for l <= r {
		if lmax <= rmax { // 左边的为瓶颈
			ans += max(0, lmax-height[l])
			lmax = max(lmax, height[l])
			l++
		} else { // 右边的为瓶颈
			ans += max(0, rmax-height[r])
			rmax = max(rmax, height[r])
			r--
		}
	}
	return ans
}
