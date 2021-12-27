package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p != nil || q != nil {
		if (p == nil && q != nil) || (p != nil && q == nil) {
			return false
		}
		if p.Val != q.Val {
			return false
		}
	}

	var leftValid bool
	if p.Left != nil || q.Left != nil {
		if (p.Left == nil && q.Left != nil) || (p.Left != nil && q.Left == nil) {
			return false
		}
		leftValid = isSameTree(p.Left, q.Left)
		if !leftValid {
			return false
		}
	}

	var rightValid bool
	if p.Right != nil || q.Right != nil {
		if (p.Right == nil && q.Right != nil) || (p.Right != nil && q.Right == nil) {
			return false
		}
		rightValid = isSameTree(p.Right, q.Right)

		if !rightValid {
			return false
		}
	}

	return true
}
