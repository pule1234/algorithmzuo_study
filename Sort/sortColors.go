package Sort

func sortColors(nums []int) {
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, l, r int) {
	if l == r {
		return
	}
	m := (l + r) / 2
	sort(nums, l, m)
	sort(nums, m+1, r)
	merge(nums, l, r, m)
}

var help [50001]int

func merge(nums []int, l, r, m int) {
	i, a, b := l, l, m+1
	for a <= m && b <= r {
		if nums[a] <= nums[b] {
			help[i] = nums[a]
			a++
			i++
		} else {
			help[i] = nums[b]
			b++
			i++
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
}
