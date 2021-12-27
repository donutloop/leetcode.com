package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	swap(root)
	return root
}

func swap(node *TreeNode) {
	if node == nil {
		return
	}

	if node.Left != nil && node.Right != nil {
		node.Left, node.Right = node.Right, node.Left
		swap(node.Right)
		swap(node.Left)
		return
	}

	if node.Left != nil {
		node.Right = node.Left
		node.Left = nil
		swap(node.Right)
		return
	}

	if node.Right != nil {
		node.Left = node.Right
		node.Right = nil
		swap(node.Left)
		return
	}

	return
}
