package width_First


package main

import (
"fmt"
"sort"
"strings"
)

const MAXN = 401

var queue [MAXN]string
var l, r int

// 下标0 -> a
// 下标1 -> b
// 下标2 -> c
// ...
// 下标25 -> z
var graph [][]string

var visited map[string]bool

func init() {
	graph = make([][]string, 26)
	for i := 0; i < 26; i++ {
		graph[i] = make([]string, 0)
	}
	visited = make(map[string]bool)
}

func minStickers(stickers []string, target string) int {
	for i := 0; i < 26; i++ {
		graph[i] = graph[i][:0]
	}
	visited = make(map[string]bool)

	for _, str := range stickers {
		str = sortString(str)
		for i := 0; i < len(str); i++ {
			if i == 0 || str[i] != str[i-1] {
				graph[str[i]-'a'] = append(graph[str[i]-'a'], str)
			}
		}
	}

	target = sortString(target)
	visited[target] = true
	l, r = 0, 0
	queue[r] = target
	r++
	level := 1

	for l < r {
		size := r - l
		for i := 0; i < size; i++ {
			cur := queue[l]
			l++
			for _, s := range graph[cur[0]-'a'] {
				next := nextString(cur, s)
				if next == "" {
					return level
				} else if !visited[next] {
					visited[next] = true
					queue[r] = next
					r++
				}
			}
		}
		level++
	}

	return -1
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func nextString(t, s string) string {
	var builder strings.Builder
	i, j := 0, 0
	for i < len(t) {
		if j == len(s) {
			builder.WriteByte(t[i])
			i++
		} else {
			if t[i] < s[j] {
				builder.WriteByte(t[i])
				i++
			} else if t[i] > s[j] {
				j++
			} else {
				i++
				j++
			}
		}
	}
	return builder.String()
}

func main() {
	stickers := []string{"ab", "ac", "bd", "de"}
	target := "abcde"
	result := minStickers(stickers, target)
	fmt.Println(result)
}
