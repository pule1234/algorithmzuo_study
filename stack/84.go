package stack

func largestRectangleArea(heights []int) int {
	stack := make([]int, 100001)
	r := 0
	ans := 0
	//找到左右比他小的值
	for i := 0; i < len(heights); i++ {
		for r > 0 && heights[stack[r-1]] >= heights[i] {
			r--
			cur := stack[r]
			var left int
			if r == 0 {
				left = -1
			} else {
				left = stack[r-1] // 当前数字压住的数字
			}
			ans = max(ans, heights[cur]*(i-left-1))
		}
		stack[r] = i
		r++
	}

	for r > 0 {
		r--
		cur := stack[r]
		var left int
		if r == 0 {
			left = -1
		} else {
			left = stack[r-1] // 当前数字压住的数字
		}
		ans = max(ans, heights[cur]*(len(heights)-left-1))
	}

	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
