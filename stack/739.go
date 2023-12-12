package stack

// 求该数最右边比他大离他最近的数
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 100001)
	r := 0
	n := len(temperatures)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		// 相等时候的处理，相等也加入单调栈
		for r > 0 && temperatures[stack[r-1]] < temperatures[i] {
			//弹出栈顶元素
			r--
			cur := stack[r]
			ans[cur] = i - cur
		}
		stack[r] = i
		r++
	}

	return ans
}
