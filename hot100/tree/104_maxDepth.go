type TreeNode struct {
	     Val int
	     Left *TreeNode
		 Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}

// BFS解法
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := make([] *TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		level_size := len(queue)
		for level_size > 0 {
			level_size--
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res++
	}
	return res
}