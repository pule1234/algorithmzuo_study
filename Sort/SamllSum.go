package Sort

// 小和问题
// 测试链接 : https://www.nowcoder.com/practice/edfe05a1d45c4ea89101d936cac32469
// 请同学们务必参考如下代码中关于输入、输出的处理
// 这是输入输出处理效率很高的写法
// 提交以下的code，提交时请把类名改成"Main"，可以直接通过

func SmallSum(arr []int) int {
	return smallsum(0, len(arr)-1, arr)
}

func smallsum(l, r int, nums []int) int {
	if l == r {
		return 0
	}
	m := (l + r) / 2
	return smallsum(l, m, nums) + smallsum(m+1, r, nums) + merge(l, r, m, nums)
}

var help [5001]int

func merge(l, r, m int, nums []int) int {
	// 首先求区间小和
	ans := 0
	for i, j, sum := l, m+1, 0; j <= r; j++ {
		sum := 0
		for i <= m && nums[i] <= nums[j] {
			sum += nums[i]
		}
		ans += sum
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

	for i := l; i <= r; i++ {
		nums[i] = help[i]
	}

	return ans
}
