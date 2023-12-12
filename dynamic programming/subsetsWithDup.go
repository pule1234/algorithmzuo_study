package dynamic_programming

import "sort"

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	f(nums, 0, make([]int, 0), 0, &ans)
	return ans
}

func f(nums []int, i int, path []int, size int, ans *[][]int) {
	if i == len(nums) {
		cur := make([]int, 0)
		for j := 0; j < len(path); j++ {
			cur = append(cur, path[j])
		}
		*ans = append(*ans, cur)
	} else {
		//下一组的第一个数的位置
		j := i + 1
		for j < len(nums) && nums[i] == nums[j] {
			j++
		}
		//当前数字要0个
		f(nums, j, path, size, ans)
		// 当前数x，要1个、要2个、要3个...都尝试
		for ; i < j; i++ {
			path = append(path, nums[i])
			size++
			f(nums, j, path, size, ans)
		}
	}
}
