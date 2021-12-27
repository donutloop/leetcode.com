package _12__Path_Sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return pathSum(root, root.Val, sum)
}

func pathSum(root *TreeNode, v int, sum int) bool {
	if root.Left == nil && root.Right == nil {
		if sum == v {
			return true
		}
		return false
	}
	var l bool
	if root.Left != nil {
		l = pathSum(root.Left, v+root.Left.Val, sum)
	}
	var r bool
	if root.Right != nil {
		r = pathSum(root.Right, v+root.Right.Val, sum)
	}
	if l || r {
		return true
	}
	return false
}
