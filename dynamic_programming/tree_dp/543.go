package tree_dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Info struct {
	diameter int
	height   int
}

func diameterOfBinaryTree(root *TreeNode) int {
	return f(root).diameter
}

func f(x *TreeNode) Info {
	if x == nil {
		return Info{0, 0}
	}

	leftinfo := f(x.Left)
	rightinfo := f(x.Right)

	height := max(leftinfo.height, rightinfo.height) + 1 // 目前这棵树上的最大值
	diamter := max(leftinfo.diameter, rightinfo.diameter)
	diamter = max(diamter, leftinfo.height+rightinfo.height)
	return Info{diamter, height}
}
