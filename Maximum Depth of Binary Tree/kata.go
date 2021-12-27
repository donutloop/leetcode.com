package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return height(root, 1)
}

func height(root *TreeNode, count int) int {
	leftHeight := count
	if root.Left != nil {
		leftHeight = height(root.Left, leftHeight+1)
	}
	rightHeight := count
	if root.Right != nil {
		rightHeight = height(root.Right, rightHeight+1)
	}
	if rightHeight > leftHeight {
		return rightHeight
	}
	return leftHeight
}
