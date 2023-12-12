package stack

// 去除重复字母保证剩余字符串的字典序最小
// 给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次
// 需保证 返回结果的字典序最小
// 要求不能打乱其他字符的相对位置
// 测试链接 : https://leetcode.cn/problems/remove-duplicate-letters/

func removeDuplicateLetters(str string) string {
	//词频表
	cnts := make([]int, 50001)
	//每种字符目前有没有进栈
	enter := make([]bool, 50001)
	stack := make([]int, 50001)
	r := 0

	s := []byte(str)
	for _, v := range s {
		cnts[v-'a']++
	}

	for _, cur := range s {
		if !enter[cur-'a'] {
			for r > 0 && stack[r-1] > int(cur) && cnts[stack[r-1]-'a'] > 0 {
				enter[stack[r-1]-'a'] = false
				r--
			}
			stack[r] = int(cur)
			r++
			enter[cur-'a'] = true
		}

		cnts[cur-'a']--
	}

	return
}
