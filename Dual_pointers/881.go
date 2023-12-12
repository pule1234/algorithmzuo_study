package Dual_pointers

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	n := len(people)
	l, r := 0, n-1
	ans := 0

	for l <= r {
		sum := 0
		if l == r {
			sum = people[l]
		} else {
			sum = people[l] + people[r]
		}

		if sum <= limit {
			l++
			r--
		} else {
			r--
		}
		ans++
	}
	return ans
}
