package _11__Minimum_Depth_of_Binary_Tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	if node.Left == nil {
		return (minDepth(node.Right) + 1)
	}
	if node.Right == nil {
		return (minDepth(node.Left) + 1)
	}
	return int(math.Min(float64(minDepth(node.Right)), float64(minDepth(node.Left)))) + 1
}
