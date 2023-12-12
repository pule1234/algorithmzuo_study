package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 字典序最小的拓扑排序
// 要求返回所有正确的拓扑排序中 字典序最小 的结果
// 建图请使用链式前向星方式，因为比赛平台用其他建图方式会卡空间
// 测试链接 : https://www.luogu.com.cn/problem/U107394

const MAXN = 100001
const MAXM = 100001

var head = make([]int, MAXN)
var next = make([]int, MAXM)
var to = make([]int, MAXM)
var cnt int

var heap = make([]int, MAXN)
var heapSize int

var indegree = make([]int, MAXM)

var ans = make([]int, MAXN)

var n, m int

func build(n int) {
	cnt = 1
	heapSize = 0
	for i := 0; i <= n; i++ {
		head[i] = 0
		indegree[i] = 0
	}
}

func addEdge(f, t int) {
	next[cnt] = head[f]
	to[cnt] = t
	head[f] = cnt
	cnt++
}

func push(num int) {
	i := heapSize
	heap[i] = num
	for heap[i] < heap[(i-1)/2] {
		swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
	heapSize++
}

func pop() int {
	ans := heap[0]
	heap[0] = heap[heapSize-1]
	heapSize--
	i := 0
	l := 1
	for l < heapSize {
		best := l
		if l+1 < heapSize && heap[l+1] < heap[l] {
			best = l + 1
		}
		if heap[best] < heap[i] {
			swap(best, i)
			i = best
			l = i*2 + 1
		} else {
			break
		}
	}
	return ans
}

func isEmpty() bool {
	return heapSize == 0
}

func swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		m, _ = strconv.Atoi(scanner.Text())
		build(n)
		for i := 0; i < m; i++ {
			scanner.Scan()
			from, _ := strconv.Atoi(scanner.Text())
			scanner.Scan()
			to, _ := strconv.Atoi(scanner.Text())
			addEdge(from, to)
			indegree[to]++
		}
		topoSort()
		for i := 0; i < n-1; i++ {
			fmt.Fprintf(writer, "%d ", ans[i])
		}
		fmt.Fprintf(writer, "%d\n", ans[n-1])
	}
}

func topoSort() {
	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			push(i)
		}
	}
	fill := 0
	for !isEmpty() {
		cur := pop()
		ans[fill] = cur
		fill++
		for ei := head[cur]; ei != 0; ei = next[ei] {
			indegree[to[ei]]--
			if indegree[to[ei]] == 0 {
				push(to[ei])
			}
		}
	}
}
