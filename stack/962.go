package stack

func maxWidthRamp(nums []int) int {
	// 令r=1相当于0位置进栈了
	// stack[0] = 0，然后栈的大小变成1
	stack := make([]int, 50001)
	r := 1

	n := len(nums)

	for i := 1; i < n; i++ {
		if nums[stack[r-1]] > nums[i] {
			stack[r] = i
			r++
		}
	}

	ans := 0
	for j := n - 1; j >= 0; j-- {
		for r > 0 && nums[stack[r-1]] <= nums[j] {
			r--
			ans = max(ans, j-stack[r])
		}
	}

	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
