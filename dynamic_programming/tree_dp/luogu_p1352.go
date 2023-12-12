package tree_dp

const MAXN = 6001

var nums [MAXN]int
var boss [MAXN]int

// 链式前向星建图
var head [MAXN]int
var next [MAXN]int
var to [MAXN]int

var cnt int
