package tree_dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	presum := make(map[int64]int)
	presum[0] = 1

	ans = 0
	f(root, targetSum, 0, presum)
	return ans
}

var ans int

// sum : 从头节点出发，来到x的时候，上方累加和是多少
// 路径必须以x作为结尾，路径累加和是target的路径数量，累加到全局变量ans上
func f(x *TreeNode, target, sum int, presum map[int64]int) {
	if x != nil {
		sum += x.Val
		ans += presum[int64(sum-target)]
		presum[int64(sum)]++
		f(x.Left, target, sum, presum)
		f(x.Right, target, sum, presum)
		// 回溯到上一层的递归
		presum[int64(sum)]--
	}
}
