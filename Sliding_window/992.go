package Sliding_window

func subarraysWithKDistinct(nums []int, k int) int {
	return numsOfMostKinds(nums, k) - numsOfMostKinds(nums, k-1)
}

const MAXN = 20001

var cnts = make([]int, MAXN)

func numsOfMostKinds(arr []int, k int) int {
	for i := 0; i <= len(arr); i++ {
		cnts[i] = 0
	}

	ans := 0
	for l, r, collect := 0, 0, 0; r < len(arr); r++ {
		if cnts[arr[r]] == 0 {
			collect++
		}
		cnts[arr[r]]++

		for collect > k {
			cnts[arr[l]]--
			if cnts[arr[l]] == 0 {
				collect--
			}
			l++
		}
		ans += r - l + 1
	}

	return ans
}
