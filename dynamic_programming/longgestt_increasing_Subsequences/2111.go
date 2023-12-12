package longgestt_increasing_Subsequences

const MAXN = 100001

var nums [MAXN]int
var ends [MAXN]int

func kIncreasing(arr []int, k int) int {
	n := len(arr)
	ans := 0

	for i := 0; i < k; i++ {
		size := 0
		for j := i; j < n; j += k {
			nums[size] = arr[j]
			size++
		}
		// 当前组长度 - 当前组最长不下降子序列长度 = 当前组至少需要修改的数字个数
		ans += size - lengthOfNoDecreasing(size)
	}
	return ans
}

// nums[0...size-1]中的最长不下降子序列长度
func lengthOfNoDecreasing(size int) int {
	length := 0
	for i := 0; i < size; i++ {
		find := bs(length, nums[i])
		if find == -1 {
			ends[length] = nums[i]
			length++
		} else {
			ends[find] = nums[i]
		}
	}
	return length
}

func bs(len, num int) int {
	l, r, m, ans := 0, len-1, 0, -1

	for l <= r {
		m = (l + r) / 2
		if num < ends[m] {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ans
}
