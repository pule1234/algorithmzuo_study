package dynamic_programming

// 字符串的全部子序列
// 子序列本身是可以有重复的，只是这个题目要求去重
// 测试链接 : https://www.nowcoder.com/practice/92e6247998294f2c933906fdedbc6e6a

func generatePermutation2(str string) []string {
	s := []byte(str)
	set := make(map[string]bool)
	f2(s, 0, make([]byte, len(s)), 0, set)
	m := len(set)
	ans := make([]string, m)
	i := 0
	for cur := range set {
		ans[i] = cur
		i++
	}
	return ans
}

func f2(s []byte, i int, path []byte, size int, set map[string]bool) {
	if i == len(s) {
		set[string(path[:size])] = true
	} else {
		path[size] = s[i]
		f2(s, i+1, path, size+1, set)
		f2(s, i+1, path, size, set)
	}
}
