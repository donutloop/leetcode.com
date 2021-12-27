package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	c := new(int)
	*c++
	count(root, c)
	return *c
}

func count(node *TreeNode, c *int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		*c++
		count(node.Left, c)
	}
	if node.Right != nil {
		*c++
		count(node.Right, c)
	}
}
