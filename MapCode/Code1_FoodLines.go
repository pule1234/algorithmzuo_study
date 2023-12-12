package main

// 最大食物链计数
// a -> b，代表a在食物链中被b捕食
// 给定一个有向无环图，返回
// 这个图中从最初级动物到最顶级捕食者的食物链有几条

// 链式前向星建图
const MAXN = 5001
const MAXM = 500001
const MOD = 80112002

var head [MAXN]int
var next [MAXM]int

var to [MAXM]int
var cnt int

var queue [MAXN]int
var indegree [MAXN]int

// 拓扑排序需要的推送信息
var lines [MAXN]int
var n, m int

func build(n int) {
	cnt = 1
	for i := 0; i < n+1; i++ {
		head[i] = 0
		lines[i] = 0
		indegree[i] = 0
	}
}

func addEdge(u, v int) {
	next[cnt] = head[u]
	to[cnt] = v
	head[u] = cnt
	cnt++
}

func ways() int {
	l, r := 0, 0
	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			queue[r] = i
			lines[i] = 1
			r++
		}
	}
	var ans int
	for l < r {
		u := queue[l]
		l++
		if head[u] == 0 {
			ans = (ans + lines[u]) % MOD
		} else {
			for ei := head[u]; ei > 0; ei = next[ei] {
				v := to[ei]
				lines[v] = (lines[v] + lines[u]) % MOD
				if indegree[v]--; indegree[v] == 0 {
					queue[r] = v
					r++
				}
			}
		}
	}
	return ans
}
