package biggest_sum

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)

	sums := make([]int, n)
	for l, r, sum := 0, 0, 0; r < n; r++ {
		sum += nums[r]
		if r-l+1 >= k {
			sums[l] = sum
			sum -= nums[l]
			l++
		}
	}

	//
	prefix := make([]int, n)
	for l, r := 1, k; r < n; l, r = l+1, r+1 {
		if sums[l] > sums[prefix[r-1]] {
			prefix[r] = l
		} else {
			prefix[r] = prefix[r-1]
		}
	}

	suffix := make([]int, n)
	suffix[n-k] = n - k
	for l := n - k - 1; l >= 0; l-- {
		if sums[l] >= sums[suffix[l+1]] {
			suffix[l] = l
		} else {
			suffix[l] = suffix[l+1]
		}
	}

	a, b, c, max := 0, 0, 0, 0
	for p, s, i, j, sum := k, k, k, 2*k-1, 0; j < n-k; i, j = i+1, j+1 {
		p = prefix[i-1]
		s = suffix[j+1]

		sum = sums[p] + sums[i] + sums[j]
		if sum > max {
			max = sum
			a = p
			b = i
			c = s
		}
	}
	return []int{a, b, c}
}
