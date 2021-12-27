package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	maxHeight := 0
	currentHeight := 0
	maxHeight = getHeight(root, maxHeight)
	return sum(root, maxHeight, currentHeight)
}

func sum(node *TreeNode, maxHeight int, currentHeight int) int {

	currentHeight++

	if currentHeight == maxHeight {
		return node.Val
	}

	var s int
	if node.Left != nil {
		s += sum(node.Left, maxHeight, currentHeight)
	}

	if node.Right != nil {
		s += sum(node.Right, maxHeight, currentHeight)
	}

	return s
}

func getHeight(node *TreeNode, h int) int {
	h++

	if node.Left == nil && node.Right == nil {
		return h
	}

	var leftHeight int
	if node.Left != nil {
		leftHeight = getHeight(node.Left, h)
	}

	var rightHeight int
	if node.Right != nil {
		rightHeight = getHeight(node.Right, h)
	}

	if leftHeight > rightHeight {
		return leftHeight
	}

	return rightHeight
}
