package Nested

import (
	"fmt"
	"sort"
	"strings"
)

// 含有嵌套的分子式求原子数量
// 测试链接 : https://leetcode.cn/problems/number-of-atoms/
func countOfAtoms(formula string) string {
	where = 0
	str := []byte(formula)
	hash := f(str, 0)
	var keys []string
	for key := range hash {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	ans := strings.Builder{}
	for _, key := range keys {
		ans.WriteString(key)
		cnt := hash[key]
		if cnt > 1 {
			ans.WriteString(fmt.Sprintf("%d", cnt))
		}
	}

	return ans.String()
}

var where int

func f(s []byte, i int) map[string]int {
	//ans 是总表
	ans := make(map[string]int)
	//之前收集到的名字，历史的一部分
	name := strings.Builder{}
	//之前收集到的有序表
	pre := make(map[string]int)
	// 历史翻几倍
	cnt := 0

	for i < len(s) && s[i] != ')' {
		if s[i] >= 'A' && s[i] <= 'Z' || s[i] == '(' {
			fill(ans, name, pre, cnt)
			name = strings.Builder{}
			pre = nil
			cnt = 0
			if s[i] >= 'A' && s[i] <= 'Z' {
				name.WriteString(string(s[i]))
				i++
			} else {
				// 遇到 (
				pre = f(s, i+1)
				i = where + 1
			}
		} else if s[i] >= 'a' && s[i] <= 'z' {
			name.WriteString(string(s[i]))
			i++
		} else {
			cnt = cnt*10 + int(s[i]-'0')
			i++
		}
	}
	fill(ans, name, pre, cnt)
	where = i
	return ans
}

func fill(ans map[string]int, name strings.Builder, pre map[string]int, cnt int) {
	if name.Len() > 0 || pre != nil { // 不为空
		if cnt == 0 {
			cnt = 1
		}
		if name.Len() > 0 {
			key := name.String()
			ans[key] = ans[key] + cnt
		} else {
			for key := range pre {
				ans[key] = ans[key] + pre[key]*cnt
			}
		}
	}
}
