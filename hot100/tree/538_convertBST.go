// 538. 把二叉搜索树转换为累加树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 一个节点的累加值等于节点值+右子树之和

func convertBST(root *TreeNode) *TreeNode {
	total :=0
	return helper(root, &total)
}

func helper(root *TreeNode, total *int) *TreeNode {
	if root == nil {
		return nil
	}
	helper(root.Right, total)
	*total = root.Val + *total
	root.Val = *total
	helper(root.Left, total)
	return root
}