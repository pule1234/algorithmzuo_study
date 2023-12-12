package main

import "strings"

// 火星词典
// 现有一种使用英语字母的火星语言
// 这门语言的字母顺序对你来说是未知的。
// 给你一个来自这种外星语言字典的字符串列表 words
// words 中的字符串已经 按这门新语言的字母顺序进行了排序 。
// 如果这种说法是错误的，并且给出的 words 不能对应任何字母的顺序，则返回 ""
// 否则，返回一个按新语言规则的 字典递增顺序 排序的独特字符串
// 如果有多个解决方案，则返回其中任意一个
// words中的单词一定都是小写英文字母组成的
// 测试链接 : https://leetcode.cn/problems/alien-dictionary/

func alineOrder(words []string) string {
	//创建入度表
	indegree := make([]int, 26)
	//等于-1代表没有这个字符出现
	for i := 0; i < len(indegree); i++ {
		indegree[i] = -1
	}

	//设置为0代表有这个字符出现
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			indegree[w[i]-'a'] = 0
		}
	}

	//创建graph
	graph := make([][]int, 0)

	for i := 0; i < len(words); i++ {
		cur := words[i]
		next := words[i+1]

		j := 0
		length := min(len(cur), len(next))
		for ; j < length; j++ {
			if cur[j] != next[j] {
				graph[cur[j]-'a'] = append(graph[cur[j]-'a'], int(next[j]-'a'))
				indegree[next[j]-'a']++
				break
			}
		}
		if j < len(cur) && j == len(next) {
			return ""
		}
	}

	queue := make([]int, 26)
	l, r := 0, 0
	kinds := 0

	for i := 0; i < 26; i++ {
		if indegree[i] != -1 {
			kinds++
		}
		if indegree[i] == 0 {
			queue[r] = i
			r++
		}
	}

	var ans strings.Builder
	for l < r {
		cur := queue[l]
		l++
		ans.WriteByte(byte(cur + 'a'))
		for _, next := range graph[cur] {
			if indegree[next]--; indegree[next] == 0 {
				queue[r] = next
				r++
			}
		}
	}

	if ans.Len() == kinds {
		return ans.String()
	} else {
		return ""
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
