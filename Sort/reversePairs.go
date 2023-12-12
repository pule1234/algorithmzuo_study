package Sort

func reversePairs(nums []int) int {
	return count(nums, 0, len(nums)-1)
}

func count(nums []int, l, r int) int {
	if l == r {
		return 0
	}
	m := (l + r) / 2
	return count(nums, l, m) + count(nums, m+1, r) + merge(nums, l, r, m)
}

var help [5001]int

func merge(nums []int, l, r, m int) int {
	ans := 0
	for i, j := l, m+1; i <= m; i++ {
		for j < len(nums) && nums[i] > 2*nums[j] {
			j++
		}
		ans += j - m - 1
	}

	a, b, i := l, m+1, l
	for a <= m && b <= r {
		if nums[a] <= nums[b] {
			help[i] = nums[a]
			i++
			a++
		} else {
			help[i] = nums[b]
			i++
			b++
		}
	}

	for a <= m {
		help[i] = nums[a]
		a++
		i++
	}

	for b <= r {
		help[i] = nums[b]
		b++
		i++
	}

	for i = l; i <= r; i++ {
		nums[i] = help[i]
	}
	return ans
}
