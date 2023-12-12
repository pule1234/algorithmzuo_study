package Nested

import "strings"

// 含有嵌套的字符串解码
// 测试链接 : https://leetcode.cn/problems/decode-string/
func decodeString(s string) string {
	where = 0
	str := []byte(s)
	return f(str, 0)
}

var where int

func f(s []byte, i int) string {
	path := strings.Builder{}
	cnt := 0
	for i < len(s) && s[i] != ']' {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			path.WriteString(string(s[i]))
			i++
		} else if s[i] >= '0' && s[i] <= '9' {
			cnt = cnt*10 + int(s[i]-'0')
			i++
		} else {
			// 遇到 [
			// cnt = 7 * ?
			path.WriteString(get(cnt, f(s, i+1)))
			i = where + 1
			cnt = 0
		}
	}
	where = i
	return path.String()
}

func get(cnt int, str string) string {
	builder := strings.Builder{}
	for i := 0; i < cnt; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
