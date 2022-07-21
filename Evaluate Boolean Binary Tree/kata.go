package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		if root.Val == 1 {
			return true
		} else {
			return false
		}
	}
	var b bool
	if root.Left != nil {
		b = evaluateTree(root.Left)
	}
	var a bool
	if root.Right != nil {
		a = evaluateTree(root.Right)
	}
	if root.Val == 2 {
		return a || b
	}
	return a && b
}
