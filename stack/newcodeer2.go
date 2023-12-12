package stack

func turns(nums []int) int {
	stack := [50001][2]int{}

	r := 0
	ans := 0
	n := len(nums)
	for i := n-1; i >=0 i--{
		cur := 0

		for  r > 0 && stack[r-1][0] < nums[i]{
			r--
			cur = max(cur+1,stack[r][1])
		}

		stack[r][0] = nums[i]
		stack[r][1] = cur
		ans = max(ans,cur)
	}
	return ans
}
