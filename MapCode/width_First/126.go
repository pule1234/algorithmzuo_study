package main

// 单词表
var dict map[string]bool
var curLevel, nextLevel map[string]bool

// 反向图
var graph map[string][]string

// 记录路径，当生成一条有效路径的时候，拷贝进ans
var path []string
var ans [][]string

func build(wordList []string) {
	dict = make(map[string]bool)
	graph = make(map[string][]string)
	ans = make([][]string, 0)
	curLevel = make(map[string]bool)
	nextLevel = make(map[string]bool)
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	build(wordList)
	if _, ok := dict[endWord]; !ok {
		return ans
	}
	if bfs(beginWord, endWord) {
		dfs(endWord, beginWord)
	}
	return ans
}

// begin -> end ，一层层bfs去，建图
// 返回值：真的能找到end，返回true；false
func bfs(begin, end string) bool {
	find := false
	curLevel[begin] = true

	for len(curLevel) > 0 {
		for k, _ := range curLevel {
			delete(dict, k)
		}

		for word, _ := range curLevel {
			// word : 去扩
			// 每个位置，字符a~z，换一遍！检查在词表中是否存在
			// 避免，加工出自己
			w := []byte(word)
			for i := 0; i < len(w); i++ {
				old := w[i]
				for ch := 'a'; ch <= 'z'; ch++ {
					w[i] = byte(ch)
					str := string(w)
					if _, ok := dict[str]; ok && str != word {
						if str == end {
							find = true
						}
						graph[str] = append(graph[str], word)
						nextLevel[str] = true
					}
				}
				w[i] = old
			}
		}
		if find {
			return true
		} else {
			tmp := curLevel
			curLevel = nextLevel
			nextLevel = tmp
			nextLevel = make(map[string]bool)
		}
	}
	return false
}

func dfs(word string, aim string) {
	path = append([]string{word}, path...)
	if word == aim {
		temp := make([]string, len(path))
		copy(temp, path)
		ans = append(ans, temp)
	} else if neighbors, ok := graph[word]; ok {
		for _, next := range neighbors {
			dfs(next, aim)
		}
	}

	path = path[:len(path)-1]
}
