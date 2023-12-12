package dynamic_programming

// 没有重复项数字的全排列
// 测试链接 : https://leetcode.cn/problems/permutations/
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	f(nums, 0, &ans)
	return ans
}

func f(nums []int, i int, ans *[][]int) {
	if i == len(nums) {
		cur := make([]int, 0)
		for _, num := range nums {
			cur = append(cur, num)
		}
		*ans = append(*ans, cur)
	} else {
		for j := i; j < len(nums); j++ {
			swap(nums, i, j)
			f(nums, i+1, ans)
			swap(nums, i, j)
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
