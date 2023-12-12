package Sort

import "math"

// 归并排序， 两种方法

func sortArray(nums []int) []int {
	mergesort1(nums)
	return nums
}

// 递归
func mergesort1(nums []int) {
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, l, r int) {
	if l == r {
		return
	}
	m := (l + r) / 2
	sort(nums, l, m)
	sort(nums, l, m+1)
	merge(nums, l, r, m)
}

// 非递归
func mergesort2(nums []int) {
	n := len(nums)

	for l, r, m, step := 1, 1, 1, 1; step < n; step <<= 1 {
		l = 0
		for l < n {
			m = (l + step) - 1
			if m+1 >= n {
				break
			}
			r = int(math.Min(float64((l+step*2)-1), float64(n-1)))
			merge(nums, l, r, m)
			l = r + 1
		}
	}
}

var help [50001]int

func merge(nums []int, l, r, m int) {
	i, a, b := l, l, m+1

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
		i++
		a++
	}

	for b <= r {
		help[i] = nums[b]
		i++
		b++
	}

	for j := l; j <= r; j++ {
		nums[j] = help[j]
	}
}
