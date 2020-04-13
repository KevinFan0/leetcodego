// 110. 平衡二叉树
import (
	"math"
)
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return root == nil || math.Abs(maxDepth(root.Left) - maxDepth(root.Right)) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return math.Max(maxDepth(root.Left),maxDepth(root.Right)) + 1
}