package queue

import (
	"sort"
)

var task []int
var worker []int
var deque [50001]int

var h, t int

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	task = tasks
	worker = workers

	sort.Ints(task)
	sort.Ints(worker)
	tsize := len(tasks)
	wsize := len(workers)
	ans := 0

	for l, r := 0, min(tsize, wsize); l <= r; {
		// m中点，一定要完成的任务数量
		m := (l + r) / 2
		if f(0, m-1, wsize-m, wsize-1, strength, pills) {
			ans = m
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return ans
}

func f(tl, tr, wl, wr int, s int, pills int) bool {
	h = 0
	t = 0
	//已经使用的药的数量
	cnt := 0

	for i, j := wl, tl; i <= wr; i++ {
		// i : 工人的编号
		// j : 任务的编号
		for ; j <= tr && task[j] <= worker[i]; j++ {
			// 工人不吃药的情况下，去解锁任务
			deque[t] = j
			t++
		}
		if h < t && task[deque[h]] <= worker[i] {
			h++
		} else {
			// 吃药之后的逻辑
			for ; j <= tr && task[j] <= worker[i]+s; j++ {
				deque[t] = j
				t++
			}

			if h < t {
				cnt++
				t--
			} else {
				return false
			}
		}

	}

	return cnt <= pills
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
