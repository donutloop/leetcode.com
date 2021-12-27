package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	set := make(map[int]struct{})
	isValid(root, set)
	if len(set) == 1 {
		return true
	}
	return false
}

func isValid(node *TreeNode, set map[int]struct{}) {
	if node == nil {
		return
	}

	set[node.Val] = struct{}{}

	if node.Left != nil {
		isValid(node.Left, set)
	}

	if node.Right != nil {
		isValid(node.Right, set)
	}
}
