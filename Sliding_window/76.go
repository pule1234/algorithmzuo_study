package Sliding_window

import (
	"math"
)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	strs := []byte(s)
	strt := []byte(t)

	//表示欠债情况表，      相当与strt中的为一个所需要的个数
	cnt := make([]int, 256)

	for _, v := range strt {
		cnt[int(v)]--
	}

	//最小的覆盖长度
	length := math.MaxInt32
	start := 0
	//start : 记录字符串开始的位置
	for l, r, debt := 0, 0, len(strt); r < len(strs); r++ {
		//debt ： 欠债的总数
		// strs[r] 当前的字符
		// cnts[s[r]] : 当前字符欠债情况，负数就是欠债，正数就是多给的
		if cnt[strs[r]]++; cnt[strs[r]] <= 0 {
			debt--
		}
		if debt == 0 { // 表示债换完了，需要判断是否可以缩短
			for cnt[strs[l]] > 0 { // 表示这种字符多了
				cnt[strs[l]]--
				l++
			}
			if r-l+1 < length {
				length = r - l + 1
				start = l
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	}

	return s[start : start+length]
}
