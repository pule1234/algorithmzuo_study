package Bidirectional_search

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]struct{})

	for _, word := range wordList {
		dict[word] = struct{}{}
	}

	if _, ok := dict[endWord]; !ok {
		return 0
	}

	//数量小的一侧
	smallLevel := make(map[string]struct{})
	//数量大的一侧
	bigLevel := make(map[string]struct{})

	//数量最小的一侧，所扩展出的下一层列表
	nextLevel := make(map[string]struct{})

	smallLevel[beginWord] = struct{}{}
	bigLevel[endWord] = struct{}{}

	for length := 2; len(smallLevel) > 0; length++ {
		for w, _ := range smallLevel {
			//从小侧开始扩展
			word := []byte(w)
			for j := 0; j < len(word); j++ {
				//每一位字符都试
				old := word[j]
				for change := 'a'; change <= 'z'; change++ {
					//每个字符从a到z都换一遍
					if byte(change) != old {
						word[j] = byte(change)
						next := string(word)
						if _, ok := bigLevel[next]; ok { // 证明重合，遍历结束
							return length
						}
						if _, ok := dict[next]; ok {
							delete(dict, next)
							nextLevel[next] = struct{}{}
						}
					}
				}
				word[j] = old
			}
		}
		if len(nextLevel) <= len(bigLevel) {
			tmp := smallLevel
			smallLevel = nextLevel
			nextLevel = tmp
		} else {
			tmp := smallLevel
			smallLevel = bigLevel
			bigLevel = nextLevel
			nextLevel = tmp
		}
		nextLevel = make(map[string]struct{})
	}

	return 0
}
