package Bellman_Ford

import "math"

// Bellman-Ford + SPFA优化模版（洛谷）
// 给定一个 n个点的有向图，请求出图中是否存在从顶点 1 出发能到达的负环
// 负环的定义是：一条边权之和为负数的回路。
// 测试链接 : https://www.luogu.com.cn/problem/P3385
// 请同学们务必参考如下代码中关于输入、输出的处理
// 这是输入输出处理效率很高的写法
// 提交以下所有代码，把主类名改成Main，可以直接通过
const MAXN = 2001
const MAXM = 6001
const MAXQ = 4000001

var head = make([]int, MAXN)
var next = make([]int, MAXM)
var to = make([]int, MAXM)
var weight = make([]int, MAXM)
var cnt int
var distance = make([]int, MAXN)

// 记录中转次数
var updateCnt = make([]int, MAXN)
var queue = make([]int, MAXQ)
var l, r int
var enter = make([]bool, MAXN)

func build(n int) {
	cnt = 1
	l = 0
	r = 0
	for i := 1; i <= n; i++ {
		head[i] = 0
		enter[i] = false
		distance[i] = math.MaxInt32
		updateCnt[i] = 0
	}
}

func addEdge(u, v, w int) {
	next[cnt] = head[u]
	to[cnt] = v
	weight[cnt] = w
	head[u] = cnt
	cnt++
}

func spfa(n int) bool {
	distance[1] = 0
	updateCnt[1]++
	queue[r] = 1
	r++
	enter[1] = true

	for l < r {
		u := queue[l]
		l++
		// 出队列
		enter[u] = false
		for ei := head[u]; ei > 0; ei = next[ei] {
			v := to[ei]
			w := weight[ei]
			if distance[u]+w < distance[v] {
				distance[v] = distance[u] + w
				if !enter[v] {
					if updateCnt[v]++; updateCnt[v] == n { // 有负环
						return true
					}
					queue[r] = v
					r++
					enter[v] = true
				}
			}
		}
	}

	return false
}
