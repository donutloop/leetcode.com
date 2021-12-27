package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {

	sum := new(int)
	findLeftLeaves(root, -1, sum)
	return *sum
}

func findLeftLeaves(node *TreeNode, kind int, sum *int) {
	if node == nil {
		return
	}
	if kind == 1 && node.Left == nil && node.Right == nil {
		*sum += node.Val
		return
	}
	if node.Left != nil {
		findLeftLeaves(node.Left, 1, sum)
	}

	if node.Right != nil {
		findLeftLeaves(node.Right, 0, sum)
	}
}
