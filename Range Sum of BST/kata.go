package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	return f(root, L, R, 0)
}

func f(node *TreeNode, L int, R int, sum int) int {
	if node.Val >= L && node.Val <= R {
		sum = sum + node.Val
	}

	if node.Left != nil && L < node.Val {
		sum = f(node.Left, L, R, sum)
	}
	if node.Right != nil && node.Val < R {
		sum = f(node.Right, L, R, sum)
	}

	return sum
}
